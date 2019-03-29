package common

import (
	"sync"
	"time"

	"github.com/iost-official/go-iost/metrics"
)

// Mode
type ModeType uint

const (
	ModeNormal ModeType = iota
	ModeSync
	ModeInit
)

var (
	mode      = ModeInit
	modeMutex = new(sync.RWMutex)
	modeGauge = metrics.NewGauge("iost_node_mode", nil)
)

// Mode return the mode of pob.
func Mode() string {
	modeMutex.RLock()
	defer modeMutex.RUnlock()

	switch mode {
	case ModeNormal:
		return "ModeNormal"
	case ModeSync:
		return "ModeSync"
	case ModeInit:
		return "ModeInit"
	default:
		return "Undefined"
	}
}

func SetMode(m ModeType) {
	modeMutex.Lock()
	defer modeMutex.Unlock()

	mode = m
	modeGauge.Set(float64(mode), nil)
}

// Witness
const (
	VoteInterval       = 1200
	SlotTime           = 3 * time.Second
	BlockNumPerWitness = 6
)

// IsWitness will judage if a public key is a witness.
func IsWitness(w string, witnessList []string) bool {
	for _, v := range witnessList {
		if v == w {
			return true
		}
	}
	return false
}

// WitnessOfNanoSec will return which witness is the current time.
func WitnessOfNanoSec(nanosec int64, witnessList []string) string {
	slot := nanosec / int64(SlotTime)
	index := slot % int64(len(witnessList))
	witness := witnessList[index]
	return witness
}

// SlotOfNanoSec will return current slot number.
func SlotOfNanoSec(nanosec int64) int64 {
	return nanosec / int64(SlotTime)
}

// TimeUntilNextSchedule will return the time left in the next slot.
func TimeUntilNextSchedule(timeSec int64) int64 {
	currentSlot := timeSec / int64(SlotTime)
	return (currentSlot+1)*int64(SlotTime) - timeSec
}
