package main

import "github.com/sschepens/pla/boomer"

// Interface determines the interface for Pla's User Interfaces
type Interface interface {
	Start(b *boomer.Boomer)
	ProcessResult(res boomer.Result)
	End()
}
