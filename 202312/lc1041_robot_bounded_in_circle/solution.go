package main

func isRobotBounded(instructions string) bool {
	x, y := 0, 0
	direction := 0
	for i := 0; i < len(instructions); i++ {
		action := instructions[i]
		switch action {
		case 'G':
			switch direction {
			case 0:
				y += 1
			case 90:
				x += 1
			case 180:
				y -= 1
			case 270:
				x -= 1
			}
		case 'L':
			direction -= 90
			direction = (360 + direction) % 360
		case 'R':
			direction += 90
			direction = (360 + direction) % 360
		}
	}
	return (x == 0 && y == 0) || ((x != 0 || y != 0) && direction != 0)
}
