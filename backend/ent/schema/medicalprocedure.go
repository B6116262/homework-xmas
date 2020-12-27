package schema

import "github.com/facebookincubator/ent"

// MedicalProcedure holds the schema definition for the MedicalProcedure entity.
type MedicalProcedure struct {
	ent.Schema
}

// Fields of the MedicalProcedure.
func (MedicalProcedure) Fields() []ent.Field {
	return nil
}

// Edges of the MedicalProcedure.
func (MedicalProcedure) Edges() []ent.Edge {
	return nil
}
