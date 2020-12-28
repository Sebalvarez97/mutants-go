package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// Trainer struct
type Dna struct {
	Id              primitive.ObjectID `json:"id" bson:"_id, gin"`
	Chain           [][]byte           `json:"chain" bson:"chain, omitempty"`
	IsMutant        bool               `json:"is_mutant" bson:"is_mutant, omitempty"`
	MutantSequences int                `json:"mutant_sequences" bson:"mutant_sequences, omitempty"`
}

// Create trainer
func NewDna(chain [][]byte, isMutant bool, sequences int) *Dna {
	return &Dna{
		Chain:           chain,
		IsMutant:        isMutant,
		MutantSequences: sequences,
	}
}
