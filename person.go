package stripe

import "encoding/json"

// RelationshipParams is used to set the relationship between an acount and a person.
type RelationshipParams struct {
	Controller       *bool    `form:"controller"`
	Director         *bool    `form:"director"`
	Email            *string  `form:"email"`
	Owner            *bool    `form:"owner"`
	PercentOwnership *float64 `form:"percent_ownership"`
	Phone            *string  `form:"phone"`
	Representative   *bool    `form:"representative"`
	Title            *string  `form:"title"`
}

// PersonParams is the set of parameters that can be used when creating or updating a person.
// For more details see https://stripe.com/docs/api#create_person.
type PersonParams struct {
	Params           `form:"*"`
	Account          *string                     `form:"-"` // Included in URL
	Address          *AccountAddressParams       `form:"address"`
	AddressKana      *AccountAddressParams       `form:"address_kana"`
	AddressKanji     *AccountAddressParams       `form:"address_kanji"`
	DOB              *DOBParams                  `form:"dob"`
	FirstName        *string                     `form:"first_name"`
	FirstNameKana    *string                     `form:"first_name_kana"`
	FirstNameKanji   *string                     `form:"first_name_kanji"`
	Gender           *string                     `form:"gender"`
	LastName         *string                     `form:"last_name"`
	LastNameKana     *string                     `form:"last_name_kana"`
	LastNameKanji    *string                     `form:"last_name_kanji"`
	MaidenName       *string                     `form:"maiden_name"`
	PersonalIDNumber *string                     `form:"personal_id_number"`
	Relationship     *RelationshipParams         `form:"relationship"`
	SSNLast4         *string                     `form:"ssn_last_4"`
}

// PersonListParams is the set of parameters that can be used when listing persons.
// For more detail see https://stripe.com/docs/api#list_persons.
type PersonListParams struct {
	ListParams `form:"*"`
	Account    *string `form:"-"` // Included in URL
	Director   *bool   `form:"director"`
	Executive  *bool   `form:"executive"`
	Owner      *bool   `form:"owner"`
}

// Relationship represents extra information needed for a Person.
type Relationship struct {
	Controller       bool    `json:"controller"`
	Director         bool    `json:"director"`
	Email            string  `json:"email"`
	Owner            bool    `json:"owner"`
	PercentOwnership float64 `json:"percent_ownership"`
	Phone            string  `json:"phone"`
	Representative   bool    `json:"representative"`
	Title            string  `json:"title"`
}

// Relationship represents the relationship between a Person and an Account.
type Requirements struct {
	CurrentlyDue  []string `json:"currently_due"`
	EventuallyDue []string `json:"eventually_due"`
	PastDue       []string `json:"past_due"`
}

// Person is the resource representing a Stripe person.
// For more details see https://stripe.com/docs/api#persons.
type Person struct {
	Account          string                `json:"account"`
	Address          *AccountAddress       `json:"address"`
	AddressKana      *AccountAddress       `json:"address_kana"`
	AddressKanji     *AccountAddress       `json:"address_kanji"`
	DOB              *DOB                  `json:"dob"`
	FirstName        string                `json:"first_name"`
	FirstNameKana    string                `json:"first_name_kana"`
	FirstNameKanji   string                `json:"first_name_kanji"`
	Gender           string                `json:"gender"`
	ID               string                `json:"id"`
	LastName         string                `json:"last_name"`
	LastNameKana     string                `json:"last_name_kana"`
	LastNameKanji    string                `json:"last_name_kanji"`
	MaidenName       string                `json:"maiden_name"`
	Relationship     *Relationship         `json:"relationship"`
	Requirements     *Requirements         `json:"requirements"`
	SSNLast4Provided bool                  `json:"ssn_last_4_provided"`
	Verification     *IdentityVerification `json:"verification"`
}

// PersonList is a list of persons as retrieved from a list endpoint.
type PersonList struct {
	ListMeta
	Data []*Person `json:"data"`
}

// UnmarshalJSON handles deserialization of a Person.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *Person) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = id
		return nil
	}

	type person Person
	var v person
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = Person(v)
	return nil
}
