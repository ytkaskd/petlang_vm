package rte

type ValueType int

const (
	Byte ValueType = iota + 1
	Integer
	Float
	Reference
)
