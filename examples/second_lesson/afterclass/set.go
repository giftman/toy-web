package main

type Set interface {
	Put(key string)
	Keys() []string
	Contains(key string) bool
	Remove(key string)
	// 如果之前已经有了，就返回旧的值，absent =false
	// 如果之前没有，就塞下去，返回 absent = true
	PutIfAbsent(key string) (old string, absent bool)
}

type SetBase struct {
	hash map[string]bool
}

var _ Set = &SetBase{}

func (s *SetBase) Put(key string) {
	s.hash[key] = true
}
func (s *SetBase) Keys() []string {
	keys := []string{}
	for k, _ := range s.hash {
		keys = append(keys, k)
	}
	return keys
}

func (s *SetBase) Contains(key string) bool {
	if _, ok := s.hash[key]; ok {
		return true
	} else {
		return false
	}
}

func (s *SetBase) Remove(key string) {
	if _, ok := s.hash[key]; ok {
		s.hash[key] = false
	}
}

func (s *SetBase) PutIfAbsent(key string) (old string, absent bool) {
	if _, ok := s.hash[key]; ok {
		return key, false
	} else {
		s.hash[key] = true
		return key, true
	}
}
func main() {

}
