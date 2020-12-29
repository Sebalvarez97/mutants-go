package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Dna struct {
	Id              primitive.ObjectID `json:"id" bson:"_id, gin"`
	DnaHash         string             `json:"dna_hash" bson:"dna_hash"`
	Chain           [][]byte           `json:"chain" bson:"chain"`
	IsMutant        bool               `json:"is_mutant" bson:"is_mutant"`
	MutantSequences int                `json:"mutant_sequences" bson:"mutant_sequences"`
}

func NewDna(chain [][]byte, hash string, isMutant bool, sequences int) *Dna {
	return &Dna{
		Id:              primitive.NewObjectID(),
		DnaHash:         hash,
		Chain:           chain,
		IsMutant:        isMutant,
		MutantSequences: sequences,
	}
}
