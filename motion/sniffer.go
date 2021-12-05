package motion

func hardBoundaryPassable(x int, y int, width int, height int, size int) bool {
	if x > width || y > height || x < 0 || y < 0 {
		return false
	}
	return true
}
