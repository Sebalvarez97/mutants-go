package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Dna struct {
	Id              primitive.ObjectID `json:"id" bson:"_id, gin"`
	DnaHash         string             `json:"dna_hash" bson:"dna_hash"`
	IsMutant        bool               `json:"is_mutant" bson:"is_mutant"`
	MutantSequences int                `json:"mutant_sequences" bson:"mutant_sequences"`
}

func NewDna(hash string, isMutant bool, sequences int) *Dna {
	return &Dna{
		Id:              primitive.NewObjectID(),
		DnaHash:         hash,
		IsMutant:        isMutant,
		MutantSequences: sequences,
	}
}
