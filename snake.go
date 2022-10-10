package main

type food struct {
	x, y int
}

type foodLocation struct {
	x, y int
}

type snake struct {
	body      []foodLocation
	length    int
	direction int
}

func (s *snake) getHead() foodLocation {
	return s.body[len(s.body)-1]
}

func (s *snake) hitWall(m model) bool {

	h := s.getHead()
	if h.x > m.height || h.y > m.width || h.x <= 0 || h.y <= 0 {
		return true
	}

	return false
}

func (s *snake) hitSelf(c foodLocation) bool {
	for _, part := range s.body {
		if part.x == c.x && part.y == c.y {
			return true
		}
	}

	return false
}
