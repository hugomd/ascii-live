package frames

import "time"

var Clock = FrameType{
	GetFrame:  getClockFrame,
	GetLength: getClockLength,
}

func getClockFrame(i int) string {
	t := time.Now().Format(time.RFC3339)
	return t
}

func getClockLength() int {
	return 0
}
