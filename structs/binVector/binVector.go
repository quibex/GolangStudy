package main

// IntSet представляет собой множество небольших неотрицательных
// целых чисел. Нулевое значение представляет пустое множество.

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

// Has указывает, содержит ли множество неотрицательное значение х.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add добавляет неотрицательное значение x в множество,
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) AddAll(all ...int) {
	for _, x := range all {
		word, bit := x/64, uint(x%64)
		for word >= len(s.words) {
			s.words = append(s.words, 0)
		}
		s.words[word] |= 1 << bit
	}
}
	
// UnionWith делает множество s равным объединению множеств s и t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
				s.words = append(s.words, tword)
		}
	}
}
	

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 { // если в этом деапазоне нет чисел  
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1 << uint64(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", i*64+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	len := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1 << uint(j)) != 0 {
				len++
			}
		}
	}
	return len
}

func (s *IntSet) Remove(x int){
	word := x/64
	bit := uint(x%64)
	if (s.words[word])&(1 << bit) != 0 {
		s.words[word] ^= (1 << bit) 
	}
}

func (s *IntSet) Copy() *IntSet {
	return s
}


func main(){
	var x, y, z IntSet 
	x.Add(1) 
	x.Add(144) 
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"
	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"
	fmt.Println(y.Len())
	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x.Len())
	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	x.Remove(144)
	fmt.Println(x.String())

	z = x
	fmt.Println(z.String())
	z.Add(555)
	x.AddAll(1, 534, 3944625637, 2742)
	fmt.Println(x.String())
	fmt.Println(z.String())
}