package main

import (
	"fmt"
	"sync"
)

// Any - The type of item in stack
type Any interface{}

// Stack - Create stack item
type Stack struct{

	// data - Stores items in stack.
	// Item type is Any(interface).
	data 	[]Any
	
	// mux - Locks and unlocks for 
	// concurrent applications in stack.
	mux		sync.RWMutex
}

// Len - Returns size of items in stack.
func (s *Stack) Len() int{

	// Lock and unlock the mutex.
	s.mux.Lock()
	defer s.mux.Unlock()
	return len(s.data)
}

// IsEmpty - Checks is the stack empty or not.
func (s *Stack) IsEmpty() bool{

	// Lock and unlock the mutex.
	s.mux.Lock()
	defer s.mux.Unlock()
	return len(s.data) == 0
}

// Push - Adds an element at the end of stack.
func (s *Stack) Push(a Any) {

	// Lock the mutex.
	s.mux.Lock()
	// Apply append operation.
	s.data = append(s.data, a)
	// Unlock the mutex.
	s.mux.Unlock()
}

// Pop - Discards the last element of stack.
func (s *Stack) Pop() Any{
	
	// Lock and Unlokc the mutex.
	s.mux.Lock()
	defer s.mux.Unlock()
	// Check if stack is empty.
	if len(s.data) == 0{
		return 	nil
	}
	// Apply discard operation.
	l := len(s.data)
	tmp := s.data[l-1]
	s.data = s.data[: l-1]
	return tmp
}

// Top - Returns the last element in stack.
func (s *Stack) Top() Any{
	// Lock the mutex.
	s.mux.Lock()
	// Unlock the mutex
	defer s.mux.Unlock()
	
	// Check if stack is empty.
	if len(s.data) == 0{
		return nil
	}
	return s.data[len(s.data)-1]
}

// Show - Displays the elements that are stored.
func (s *Stack) Show() Any{
	// Lock the mutex.
	s.mux.Lock()
	// Unlock the mutex
	defer s.mux.Unlock()
	
	// Check if stack is empty.
	if len(s.data) == 0{
		return nil
	}
	return s.data
}


func main() {
  
  // Create new Stack struct.
	S := new(Stack)
	fmt.Println("empty? ", S.IsEmpty())
  
  // Push elements.
	S.Push("cat")
	S.Push(1999)
	S.Push("Car")
	S.Push(11.111)
	S.Push(true)
  
  // Print.
	fmt.Println(S.Show())
	fmt.Println("len(S) ", S.Len())
	fmt.Println("pop ", S.Pop())
	fmt.Println("pop ", S.Pop())
	fmt.Println(S.Show())
	fmt.Println("len(S) ", S.Len())
	fmt.Println("top ", S.Top())
	fmt.Println("empty? ", S.IsEmpty())
  
}
