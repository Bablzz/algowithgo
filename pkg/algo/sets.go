package algo

type Set struct {
	stringMap map[string]bool
}

func (s *Set) New() {
	s.stringMap = make(map[string]bool)
}

func (s *Set) Add(element string) {
	if !s.isContains(element) {
		s.stringMap[element] = true
	}
}

func (s *Set) isContains(element string) bool {
	var exists bool
	_, exists = s.stringMap[element]
	return exists
}

func (s *Set) Delete(element string) {
	delete(s.stringMap, element)
}

func (s *Set) Get(element string) (string, bool) {
	for elem := range s.stringMap {
		if elem == element {
			return elem, true
		}
	}
	return "", false
}
