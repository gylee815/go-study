package slot

import (
	"math/rand"
	"time"
)

type RESULT int

const (
	Jackpot RESULT = iota
	Equalone
	Equalzero
)

func InitSlotMachine(m *[3][10]int) {
	for i := range *m {
		for j := range m[i] {
			m[i][j] = j + 1
		}
	}
}

func RunSlotMachine(m *[3][10]int) [3]int {
	rand.Seed(time.Now().UnixNano())
	var slot [3]int

	for i := range slot {
		randIdx := rand.Intn(10)
		slot[i] = m[i][randIdx]
	}

	return slot
}

// func SlotResult(slot *[3]int) *[3]int {
// 	sort.Ints(slot[:])
// 	return slot
// }

func SlotSort(slot *[3]int) RESULT {
	for i := range slot {
		if i+1 == len(slot) {
			break
		}
		idx := i
		tmp := slot[idx+1]
		for idx >= 0 && tmp < slot[idx] {
			slot[idx+1] = slot[idx]
			idx--
			// fmt.Println(*slot)
		}
		slot[idx+1] = tmp
		// fmt.Println(*slot)
	}

	switch true {
	case slot[0] == slot[1] && slot[1] == slot[2]:
		return Jackpot
	case slot[0] == slot[1] || slot[1] == slot[2] || slot[0] == slot[2]:
		return Equalone
	case slot[0] != slot[1] || slot[1] != slot[2] || slot[0] != slot[2]:
		return Equalzero
	default:
		return Equalzero
	}

}
