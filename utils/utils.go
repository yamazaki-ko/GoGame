package utils

// nowPosition : 現在の位置
var nowPosition int

// GetPosition : Get now position
func GetPosition() int {
	return nowPosition
}

// SetPosition : Set now position
func SetPosition(d int) {
	nowPosition = d
}

// Contains is check
func Contains(s []int, e int) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
