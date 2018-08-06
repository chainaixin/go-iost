package pob

import (
	"fmt"
	"github.com/iost-official/Go-IOS-Protocol/core/new_blockcache"
	"github.com/iost-official/Go-IOS-Protocol/db"
	"time"
	"github.com/iost-official/Go-IOS-Protocol/core/new_txpool"
	"encoding/binary"
	"github.com/iost-official/Go-IOS-Protocol/core/new_block"
	. "github.com/iost-official/Go-IOS-Protocol/account"
	. "github.com/iost-official/Go-IOS-Protocol/core/new_tx"
	. "github.com/iost-official/Go-IOS-Protocol/new_consensus/common"
	"github.com/iost-official/Go-IOS-Protocol/common"
	"errors"
)

func genGenesis(initTime int64) *block.Block {
	var code string
	for k, v := range GenesisAccount {
		code += fmt.Sprintf("@PutHM iost %v f%v\n", k, v)
	}

	tx := Tx{
		Time: 0,
		// TODO what is the genesis tx?
	}

	genesis := &block.Block{
		Head: block.BlockHead{
			Version: 0,
			Number:  0,
			Time:    initTime,
		},
		Txs:      make([]Tx, 0),
		Receipts: make([]TxReceipt, 0),
	}
	genesis.Txs = append(genesis.Txs, tx)
	return genesis
}

func genBlock(acc Account, node *blockcache.BlockCacheNode, db *db.MVCCDB) *block.Block {
	lastBlk := node.Block
	blk := block.Block{
		Head: block.BlockHead{
			Version:    0,
			ParentHash: lastBlk.HeadHash(),
			Number:     lastBlk.Head.Number + 1,
			Witness:    acc.ID,
			Time:       GetCurrentTimestamp().Slot,
		},
		Txs:      []Tx{},
		Receipts: []TxReceipt{},
	}

	txCnt := 1000

	limitTime := time.NewTicker(((SlotLength/3 - 1) + 1) * time.Second)
	if new_txpool.TxPoolS != nil {
		tx, err := new_txpool.TxPoolS.PendingTransactions(txCnt)
		if err == nil {
			txPoolSize.Set(float64(new_txpool.TxPoolS.TransactionNum()))

			if len(tx) != 0 {
				VerifyTxBegin(lastBlk, db)
			ForEnd:
				for _, t := range tx {
					select {
					case <-limitTime.C:
						break ForEnd
					default:
						if len(blk.Txs) >= txCnt {
							break ForEnd
						}
						if receipt, err := VerifyTx(t); err == nil {
							db.Commit()
							blk.Txs = append(blk.Txs, *t)
							blk.Receipts = append(blk.Receipts, receipt)
						} else {
							db.Rollback()
						}
					}
				}
			}
		}
	}

	blk.Head.TxsHash = blk.CalculateTxsHash()
	blk.Head.MerkleHash = blk.CalculateMerkleHash()
	headInfo := generateHeadInfo(blk.Head)
	sig, _ := common.Sign(common.Secp256k1, headInfo, acc.Seckey)
	blk.Head.Signature = sig.Encode()
	db.Tag(blk.HeadHash())

	generatedBlockCount.Inc()

	return &blk
}

func generateHeadInfo(head block.BlockHead) []byte {
	var info, numberInfo, versionInfo []byte
	info = make([]byte, 8)
	versionInfo = make([]byte, 4)
	numberInfo = make([]byte, 4)
	binary.BigEndian.PutUint64(info, uint64(head.Time))
	binary.BigEndian.PutUint32(versionInfo, uint32(head.Version))
	binary.BigEndian.PutUint32(numberInfo, uint32(head.Number))
	info = append(info, versionInfo...)
	info = append(info, numberInfo...)
	info = append(info, head.ParentHash...)
	info = append(info, head.TxsHash...)
	info = append(info, head.MerkleHash...)
	info = append(info, head.Info...)
	return common.Sha256(info)
}

func verifyBasics(blk *block.Block) error {
	// verify block witness
	if witnessOfTime(Timestamp{Slot: blk.Head.Time}) != blk.Head.Witness {
		return errors.New("wrong witness")
	}

	headInfo := generateHeadInfo(blk.Head)
	var signature common.Signature
	signature.Decode(blk.Head.Signature)

	if blk.Head.Witness != common.Base58Encode(signature.Pubkey) {
		return errors.New("wrong pubkey")
	}

	// verify block witness signature
	if !common.VerifySignature(headInfo, signature) {
		return errors.New("wrong signature")
	}

	// block produced by itself: do not verify the rest parts
	if blk.Head.Witness == staticProp.ID {
		return nil
	}

	// verify slot map
	if staticProp.hasSlotWitness(uint64(blk.Head.Time), blk.Head.Witness) {
		return errors.New("witness slot duplicate")
	}

	return nil
}

func verifyBlockHead(blk *block.Block, parent *block.Block) error {
	return VerifyBlockHead(blk, parent)
}

func verifyBlockTxs(blk *block.Block, db *db.MVCCDB) error {
	// verify exist txs
	for _, tx := range blk.Txs {
		if blk.Head.Time - tx.Time > 300 {
			return errors.New("tx too old")
		}
		exist := new_txpool.TxPoolS.ExistTx(tx.Hash(), blk)
		if exist == INBLOCK {
			return errors.New("duplicate tx")
		} else if exist != PENDING {
			if err := tx.VerifySelf(); err != nil {
				return errors.New("tx wrong signature")
			}
		}
	}

	// verify txs
	err := VerifyBlock(blk, db)
	if err != nil {
		return err
	}
	return nil
}

func updateNodeInfo(node *blockcache.BlockCacheNode) {
	node.Number = uint64(node.Block.Head.Number)
	node.Witness = node.Block.Head.Witness

	// watermark
	node.ConfirmUntil = staticProp.Watermark[node.Witness]
	staticProp.Watermark[node.Witness] = node.Number + 1

	// slot map
	staticProp.addSlotWitness(uint64(node.Block.Head.Time), node.Witness)
}

func updatePendingWitness(node *blockcache.BlockCacheNode, db *db.MVCCDB) []string {
	// pending witness
	newList, err := db.Get("state", "witnessList")

	if err == nil {
		node.PendingWitnessList = newList
		node.LastWitnessListNumber = node.Number
	} else {
		node.PendingWitnessList = node.Parent.PendingWitnessList
		node.LastWitnessListNumber = node.Parent.LastWitnessListNumber
	}
	return nil
}

func calculateConfirm(node *blockcache.BlockCacheNode, root *blockcache.BlockCacheNode) *blockcache.BlockCacheNode {
	// return the last node that confirmed
	confirmNumber := staticProp.NumberOfWitnesses*2/3 + 1
	startNumber := node.Number
	topNumber := root.Number
	confirmMap := make(map[string]int)
	confirmUntil := make([][]string, startNumber-topNumber + 1)
	for node != root {
		if node.ConfirmUntil <= node.Number {
			if num, err := confirmMap[node.Witness]; err {
				confirmMap[node.Witness] = 1
			} else {
				confirmMap[node.Witness] = num + 1
			}
		}
		index := int64(node.ConfirmUntil) - int64(topNumber)
		if index > 0 {
			confirmUntil[index] = append(confirmUntil[index], node.Witness)
		}
		if len(confirmMap) >= confirmNumber {
			staticProp.delSlotWitness(topNumber, node.Number)
			return node
		}
		i := node.Number - topNumber
		if confirmUntil[i] != nil {
			for _, witness := range confirmUntil[i] {
				confirmMap[witness]--
				if confirmMap[witness] == 0 {
					delete(confirmMap, witness)
				}
			}
		}
		node = node.Parent
	}
	return nil
}

func promoteWitness(node *blockcache.BlockCacheNode, confirmed *blockcache.BlockCacheNode) {
	// update the last pending witness list that has been confirmed
	for node != confirmed && node.LastWitnessListNumber > confirmed.Number {
		node = node.Parent
	}
	if node.PendingWitnessList != nil {
		staticProp.updateWitnessList(node.PendingWitnessList)
	}
}
