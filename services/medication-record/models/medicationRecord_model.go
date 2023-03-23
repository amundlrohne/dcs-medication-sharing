package models

type Consent struct {
	ConsentID string `json:"consentid"`
}

type Identifier struct {
	Value string `json:"value"`
}

type Dosage struct {
	Sequence int    `json:"sequence"`
	Text     string `json:"text"`
}

type Note struct {
	Text string `json:"text"`
}

type MedicationCodeableConcept struct {
	Text string `json:"text"`
}

type Subject struct {
	Display string `json:"display"`
}

type MedicationStatement struct {
	ResourceType              string                    `json:"resourceType"`
	ID                        string                    `json:"id"`
	Subject                   Subject                   `json:"subject"`
	Status                    string                    `json:"status"`
	MedicationCodeableConcept MedicationCodeableConcept `json:"medicationCodeableConcept"`
	Note                      [1]Note                   `json:"note"`
	EffectiveDateTime         string                    `json:"effectiveDateTime"`
	Dosage                    [1]Dosage                 `json:"dosage"`
	Identifier                [1]Identifier             `json:"identifier"`
}

type Resource struct {
	MedicationStatement MedicationStatement `json:"resource"`
}

type Bundle struct {
	ResourceType string     `json:"resourceType"`
	ID           string     `json:"id"`
	Identifier   Identifier `json:"identifier"`
	Type         string     `json:"type"`
	Entry        []Resource `json:"entry"`
}

type MedicationRecord struct {
	ConsentID         string `json:"consentid,omitempty"`                             //Consent ID
	Name              string `json:"name,omitempty" validate:"required"`              // Patient Name
	Medication        string `json:"medication,omitempty"  validate:"required"`       // Medication Name
	Note              string `json:"note,omitempty" validate:"required"`              // Doctors Note
	Status            string `json:"status,omitempty" validate:"required"`            // Status of patient taking the medication
	EffectiveDateTime string `json:"effectivedatetime,omitempty" validate:"required"` // Date of when pill was first prescribed
	DosageSequence    int    `json:"dosagesequence,omitempty" validate:"required"`    // Sequence of how many times patient takes medication
	DosageNote        string `json:"dosagenote,omitempty" validate:"required"`        // Note specifying dosage
}

type MedicationRecords struct {
	Records []MedicationRecord `json:"records"`
}
