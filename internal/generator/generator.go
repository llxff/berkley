package generator

import (
	"github.com/brianvoe/gofakeit/v7"
)

type Pair struct {
	Key   string
	Value string
}

func GeneratePair(length int) *Pair {
	return &Pair{
		Key:   gofakeit.Email(),
		Value: gofakeit.Sentence(length),
	}
}
