package models

// Hash represents a collection of pointers to Node
type Hash struct {
	Capacity int
	Array    []*Node
}
