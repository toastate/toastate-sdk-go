package toastcloud

import (
	"fmt"

	"github.com/toastate/toastate-sdk-go/common/models"
)

type SignupInput struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type SignupOutput struct {
	User *models.User `json:"user,omitempty"`
}

type signupResponse struct {
	Success bool         `json:"success"`
	User    *models.User `json:"user,omitempty"`
}

func (sess *Session) Signup(req *SignupInput) (*SignupOutput, error) {
	resp := &signupResponse{}

	apierr, err := sess.client.Post("/signup", req, resp)
	if err != nil {
		return nil, err
	}
	if apierr != nil {
		return nil, fmt.Errorf("APIERROR: status: %v; code: %v; message: %v", apierr.Status, apierr.Code, apierr.Message)
	}

	if !resp.Success {
		return nil, fmt.Errorf("The API returned a failure with a 200 HTTP status code which should not happen")
	}

	if resp.User == nil {
		return nil, fmt.Errorf("The request was successfull but the remote API returned an empty body")
	}

	return &SignupOutput{
		User: resp.User,
	}, nil
}

type SigninInput struct {
	Email           string `json:"email,omitempty"`
	Password        string `json:"password,omitempty"`
	ExtendedSession bool   `json:"extended_session,omitempty"`
}

type SigninOutput struct {
	Token string `json:"token"`
}

type signinRequest struct {
	Email           string `json:"email,omitempty"`
	Password        string `json:"password,omitempty"`
	ExtendedSession bool   `json:"extended_session,omitempty"`
	SetCookie       bool   `json:"set_cookie,omitempty"`
	SetToken        bool   `json:"set_token,omitempty"`
}

type signinResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}

func (sess *Session) Signin(req *SigninInput) (*SigninOutput, error) {
	resp := &signinResponse{}

	r := &signinRequest{
		Email:           req.Email,
		Password:        req.Password,
		ExtendedSession: req.ExtendedSession,
		SetCookie:       false,
		SetToken:        true,
	}

	apierr, err := sess.client.Post("/signin", r, resp)
	if err != nil {
		return nil, err
	}
	if apierr != nil {
		return nil, fmt.Errorf("APIERROR: status: %v; code: %v; message: %v", apierr.Status, apierr.Code, apierr.Message)
	}
	if !resp.Success {
		return nil, fmt.Errorf("The API returned a failure with a 200 HTTP status code which should not happen")
	}

	if resp.Token == "" {
		return nil, fmt.Errorf("The request was successfull but the remote API did not return an authentication token")
	}

	return &SigninOutput{
		Token: resp.Token,
	}, nil
}

type SetupBillingInput struct {
}

type SetupBillingOutput struct {
	URL string `json:"url,omitempty"`
}

type setupBillingResponse struct {
	Success bool   `json:"success"`
	URL     string `json:"url,omitempty"`
}

func (sess *Session) SetupBilling(req *SetupBillingInput) (*SetupBillingOutput, error) {
	resp := &setupBillingResponse{}

	apierr, err := sess.client.AuthedPost("/user/setupbilling", nil, resp)
	if err != nil {
		return nil, err
	}
	if apierr != nil {
		return nil, fmt.Errorf("APIERROR: status: %v; code: %v; message: %v", apierr.Status, apierr.Code, apierr.Message)
	}

	if !resp.Success {
		return nil, fmt.Errorf("The API returned a failure with a 200 HTTP status code which should not happen")
	}

	if resp.URL == "" {
		return nil, fmt.Errorf("The request was successfull but the remote API returned an empty body")
	}

	return &SetupBillingOutput{
		URL: resp.URL,
	}, nil
}
