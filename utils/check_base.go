package utils

import (
	"fmt"
	"os"
)

type Status int

const (
	OK Status = iota
	WARNING
	CRITICAL
	UNKNOWN
)

type Checkbase struct {
	Message string
	Status  Status
}

func Init(st Status, msg string) *Checkbase {
	return &Checkbase{
		Message: msg,
		Status:  st,
	}
}

func (ckb *Checkbase) Exit() {
	fmt.Println(ckb.Message)
	os.Exit(int(ckb.Status))
}

func Ok(msg string) *Checkbase {
	return Init(OK, msg)
}

func Warning(msg string) *Checkbase {
	return Init(WARNING, msg)
}

func Critical(msg string) *Checkbase {
	return Init(CRITICAL, msg)
}

func Unknown(msg string) *Checkbase {
	return Init(UNKNOWN, msg)
}
