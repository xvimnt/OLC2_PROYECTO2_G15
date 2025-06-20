package interpreter

import (
	"strings"
	"sync"
)

type NativePrintSingleton struct {
	builder strings.Builder
}

var instance *NativePrintSingleton
var once sync.Once

func GetNativePrintSingleton() *NativePrintSingleton {
	once.Do(func() {
		instance = &NativePrintSingleton{}
	})
	return instance
}

func (s *NativePrintSingleton) Print(str string) {
	s.builder.WriteString(str)
}

func (s *NativePrintSingleton) Println(str string) {
	s.builder.WriteString(str)
	s.builder.WriteString("\n")
}

func (s *NativePrintSingleton) GetAndClear() string {
	defer s.builder.Reset()
	return s.builder.String()
}
