package models

type CustomDomain struct {
	RootDomain        string            `json:"root_domain,omitempty" bson:"root_domain,omitempty"`
	Domains           []string          `json:"domains,omitempty" bson:"sub_domain,omitempty"`
	ID                string            `json:"id,omitempty" bson:"id,omitempty"`
	UserID            string            `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Enabled           bool              `json:"enabled" bson:"enabled"`
	SSL               bool              `json:"ssl" bson:"ssl"`
	SSLError          string            `json:"ssl_error,omitempty" bson:"ssl_error"`
	VerificationToken string            `json:"verification_token,omitempty" bson:"verification_token,omitempty"`
	LinkedToaster     map[string]string `json:"linked_toasters,omitempty" bson:"linked_toasters"`
}
