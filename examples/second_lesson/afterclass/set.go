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

type S struct {
	hash map[string]bool
}

func (s *S) Put(key string) {
	s.hash[key] = true
}
func (s *S) Keys() []string {
	keys := []string{}
	for k, _ := range s.hash {
		keys = append(keys, k)
	}
	return keys
}

func (s *S) Contains(key string) bool {
	if _, ok := s.hash[key]; ok {
		return true
	} else {
		return false
	}
}

func (s *S) Remove(key string) {
	if _, ok := s.hash[key]; ok {
		s.hash[key] = false
	}
}

func (s *S) PutIfAbsent(key string) (old string, absent bool) {
	if _, ok := s.hash[key]; ok {
		return key, false
	} else {
		s.hash[key] = true
		return key, true
	}
}
func main() {

}
