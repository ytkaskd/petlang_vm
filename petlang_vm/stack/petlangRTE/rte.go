package rte

type ValueType int

const (
	Byte ValueType = iota + 1
	Integer
	Float
	Reference
)

type VmRTE interface{}

type PetlangBool struct {
	basis VmRTE
	Value bool
}
type PetlangByte struct {
	basis VmRTE
	Value byte
}
type PetlangInt struct {
	basis VmRTE
	Value int
}
type PetlangFLoat32 struct {
	basis VmRTE
	Value float32
}
type PetlangRef struct {
	basis VmRTE
	Value int
}
