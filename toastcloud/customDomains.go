package toastcloud

import (
	"fmt"

	"github.com/toastate/toastate-sdk-go/common/models"
)

type CreateCustomDomainInput struct {
	RootDomain    string            `json:"root_domain,omitempty"`
	Domains       []string          `json:"domains,omitempty"`
	LinkedToaster map[string]string `json:"linked_toasters,omitempty" bson:"linked_toasters"`
}

type CreateCustomDomainOutput struct {
	CustomDomain                  *models.CustomDomain `json:"custom_domain,omitempty"`
	OwnershiptCheckTXTRecordName  string               `json:"ownership_check_txt_record_name,omitempty"`
	OwnershiptCheckTXTRecordValue string               `json:"ownership_check_txt_record_value,omitempty"`
	CNAMESRecord                  map[string]string    `json:"cnames_record,omitempty"`
	Verified                      bool                 `json:"verified,omitempty"`
}

type createCustomDomainResponse struct {
	Success                       bool                 `json:"success"`
	CustomDomain                  *models.CustomDomain `json:"custom_domain,omitempty"`
	OwnershiptCheckTXTRecordName  string               `json:"ownership_check_txt_record_name,omitempty"`
	OwnershiptCheckTXTRecordValue string               `json:"ownership_check_txt_record_value,omitempty"`
	CNAMESRecord                  map[string]string    `json:"cnames_record,omitempty"`
}

func (sess *Session) CreateCustomDomain(input *CreateCustomDomainInput) (*CreateCustomDomainOutput, error) {
	resp := &createCustomDomainResponse{}

	apierr, err := sess.client.AuthedPost("/customdomain", input, resp)
	if err != nil {
		return nil, err
	}
	if apierr != nil {
		return nil, fmt.Errorf("APIERROR: status: %v; code: %v; message: %v", apierr.Status, apierr.Code, apierr.Message)
	}

	if !resp.Success {
		return nil, fmt.Errorf("The API returned a failure with a 200 HTTP status code which should not happen")
	}

	return &CreateCustomDomainOutput{
		CustomDomain:                  resp.CustomDomain,
		OwnershiptCheckTXTRecordName:  resp.OwnershiptCheckTXTRecordName,
		OwnershiptCheckTXTRecordValue: resp.OwnershiptCheckTXTRecordValue,
		CNAMESRecord:                  resp.CNAMESRecord,
	}, nil
}

type VerifyCustomDomainInput struct {
	ID string `json:"id,omitempty"`
}

type VerifyCustomDomainOutput struct {
	CustomDomain *models.CustomDomain `json:"custom_domain,omitempty"`
}

type verifyCustomDomainResponse struct {
	Success      bool                 `json:"success"`
	CustomDomain *models.CustomDomain `json:"custom_domain,omitempty"`
}

func (sess *Session) VerifyCustomDomain(input *VerifyCustomDomainInput) (*VerifyCustomDomainOutput, error) {
	if input.ID == "" {
		return nil, fmt.Errorf("you did not provide the ID of the custom domain to update")
	}

	resp := &verifyCustomDomainResponse{}

	apierr, err := sess.client.AuthedPost("/customdomain/verify/"+input.ID, nil, resp)
	if err != nil {
		return nil, err
	}
	if apierr != nil {
		return nil, fmt.Errorf("APIERROR: status: %v; code: %v; message: %v", apierr.Status, apierr.Code, apierr.Message)
	}

	if !resp.Success {
		return nil, fmt.Errorf("The API returned a failure with a 200 HTTP status code which should not happen")
	}

	return &VerifyCustomDomainOutput{
		CustomDomain: resp.CustomDomain,
	}, nil
}

type UpdateCustomDomainInput struct {
	ID string `json:"id,omitempty"`

	Domains       []string          `json:"domains,omitempty"`
	LinkedToaster map[string]string `json:"linked_toasters,omitempty" bson:"linked_toasters"`
}

type UpdateCustomDomainRequest struct {
	Domains       []string          `json:"domains,omitempty"`
	LinkedToaster map[string]string `json:"linked_toasters,omitempty" bson:"linked_toasters"`
}

type UpdateCustomDomainOutput struct {
	CustomDomain                  *models.CustomDomain `json:"custom_domain,omitempty"`
	OwnershiptCheckTXTRecordName  string               `json:"ownership_check_txt_record_name,omitempty"`
	OwnershiptCheckTXTRecordValue string               `json:"ownership_check_txt_record_value,omitempty"`
	CNAMESRecord                  map[string]string    `json:"cnames_record,omitempty"`
}

type updateCustomDomainResponse struct {
	Success                       bool                 `json:"success"`
	CustomDomain                  *models.CustomDomain `json:"custom_domain,omitempty"`
	OwnershiptCheckTXTRecordName  string               `json:"ownership_check_txt_record_name,omitempty"`
	OwnershiptCheckTXTRecordValue string               `json:"ownership_check_txt_record_value,omitempty"`
	CNAMESRecord                  map[string]string    `json:"cnames_record,omitempty"`
}

func (sess *Session) UpdateCustomDomain(input *UpdateCustomDomainInput) (*UpdateCustomDomainOutput, error) {
	if input.ID == "" {
		return nil, fmt.Errorf("you did not provide the ID of the custom domain to update")
	}

	resp := &updateCustomDomainResponse{}
	req := &UpdateCustomDomainRequest{
		Domains:       input.Domains,
		LinkedToaster: input.LinkedToaster,
	}

	apierr, err := sess.client.AuthedPut("/customdomain/"+input.ID, req, resp)
	if err != nil {
		return nil, err
	}
	if apierr != nil {
		return nil, fmt.Errorf("APIERROR: status: %v; code: %v; message: %v", apierr.Status, apierr.Code, apierr.Message)
	}

	if !resp.Success {
		return nil, fmt.Errorf("The API returned a failure with a 200 HTTP status code which should not happen")
	}

	return &UpdateCustomDomainOutput{
		CustomDomain:                  resp.CustomDomain,
		OwnershiptCheckTXTRecordName:  resp.OwnershiptCheckTXTRecordName,
		OwnershiptCheckTXTRecordValue: resp.OwnershiptCheckTXTRecordValue,
		CNAMESRecord:                  resp.CNAMESRecord,
	}, nil
}

type ListCustomDomainsInput struct {
}

type ListCustomDomainsOutput struct {
	CustomDomains []models.CustomDomain `json:"custom_domains,omitempty"`
}

type listCustomDomainsResponse struct {
	Success       bool                  `json:"success"`
	CustomDomains []models.CustomDomain `json:"custom_domains,omitempty"`
}

func (sess *Session) ListCustomDomains(input *ListCustomDomainsInput) (*ListCustomDomainsOutput, error) {
	resp := &listCustomDomainsResponse{}

	apierr, err := sess.client.AuthedGet("/customdomain/list", resp)
	if err != nil {
		return nil, err
	}
	if apierr != nil {
		return nil, fmt.Errorf("APIERROR: status: %v; code: %v; message: %v", apierr.Status, apierr.Code, apierr.Message)
	}

	if !resp.Success {
		return nil, fmt.Errorf("The API returned a failure with a 200 HTTP status code which should not happen")
	}

	return &ListCustomDomainsOutput{
		CustomDomains: resp.CustomDomains,
	}, nil
}

type GetCustomDomainInput struct {
	ID string `json:"id,omitempty"`
}

type GetCustomDomainOutput struct {
	CustomDomain                  *models.CustomDomain `json:"custom_domain,omitempty"`
	OwnershiptCheckTXTRecordName  string               `json:"ownership_check_txt_record_name,omitempty"`
	OwnershiptCheckTXTRecordValue string               `json:"ownership_check_txt_record_value,omitempty"`
	CNAMESRecord                  map[string]string    `json:"cnames_record,omitempty"`
}

type getCustomDomainResponse struct {
	Success                       bool                 `json:"success"`
	CustomDomain                  *models.CustomDomain `json:"custom_domain,omitempty"`
	OwnershiptCheckTXTRecordName  string               `json:"ownership_check_txt_record_name,omitempty"`
	OwnershiptCheckTXTRecordValue string               `json:"ownership_check_txt_record_value,omitempty"`
	CNAMESRecord                  map[string]string    `json:"cnames_record,omitempty"`
}

func (sess *Session) GetCustomDomain(input *GetCustomDomainInput) (*GetCustomDomainOutput, error) {
	if input.ID == "" {
		return nil, fmt.Errorf("you did not provide the ID of the custom domain to get")
	}

	resp := &getCustomDomainResponse{}

	apierr, err := sess.client.AuthedGet("/customdomain/"+input.ID, resp)
	if err != nil {
		return nil, err
	}
	if apierr != nil {
		return nil, fmt.Errorf("APIERROR: status: %v; code: %v; message: %v", apierr.Status, apierr.Code, apierr.Message)
	}

	if !resp.Success {
		return nil, fmt.Errorf("The API returned a failure with a 200 HTTP status code which should not happen")
	}

	return &GetCustomDomainOutput{
		CustomDomain:                  resp.CustomDomain,
		OwnershiptCheckTXTRecordName:  resp.OwnershiptCheckTXTRecordName,
		OwnershiptCheckTXTRecordValue: resp.OwnershiptCheckTXTRecordValue,
		CNAMESRecord:                  resp.CNAMESRecord,
	}, nil
}

type DeleteCustomDomainInput struct {
	ID string `json:"id,omitempty"`
}

type DeleteCustomDomainOutput struct {
}

type deleteCustomDomainResponse struct {
	Success bool `json:"success"`
}

func (sess *Session) DeleteCustomDomain(input *DeleteCustomDomainInput) (*DeleteCustomDomainOutput, error) {
	if input.ID == "" {
		return nil, fmt.Errorf("you did not provide the ID of the custom domain to delete")
	}

	resp := &deleteCustomDomainResponse{}

	apierr, err := sess.client.AuthedDelete("/customdomain/"+input.ID, nil, resp)
	if err != nil {
		return nil, err
	}
	if apierr != nil {
		return nil, fmt.Errorf("APIERROR: status: %v; code: %v; message: %v", apierr.Status, apierr.Code, apierr.Message)
	}

	if !resp.Success {
		return nil, fmt.Errorf("The API returned a failure with a 200 HTTP status code which should not happen")
	}

	return &DeleteCustomDomainOutput{}, nil
}
