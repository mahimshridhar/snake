package main

type food struct {
	x, y int
}

type coord struct {
	x, y int
}

type snake struct {
	body      []coord
	length    int
	direction int
}

func (s *snake) getHead() coord {
	return s.body[len(s.body)-1]
}

func (s *snake) hitWall(m Model) bool {

	h := s.getHead()
	if h.x >= m.height || h.y > m.width-1 || h.x <= 0 || h.y <= 0 {
		return true
	}

	return false
}

func (s *snake) hitSelf(c coord) bool {
	for _, part := range s.body {
		if part.x == c.x && part.y == c.y {
			return true
		}
	}

	return false
}
