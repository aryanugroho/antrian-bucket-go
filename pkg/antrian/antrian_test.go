package antrian

import "testing"

func TestAntrian_Start(t *testing.T) {
	tests := []struct {
		slot         Antrian
		queueProcess []int
		expected     bool
	}{
		{
			slot: Antrian{
				Slot: 3,
			},
			queueProcess: []int{2, 1, 1, 2, 2},
			expected:     true,
		},
	}

	for _, v := range tests {
		v.slot.Start(v.queueProcess)
		v.slot.Close()
	}
}
