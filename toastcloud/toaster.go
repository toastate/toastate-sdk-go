package toastcloud

import (
	"fmt"
	"io"

	"github.com/toastate/toastate-sdk-go/common/models"
	"github.com/toastate/toastate-sdk-go/internal/apiclient"
)

type ToasterCountInput struct {
	ID string `json:"id,omitempty"`
}

type ToasterCountOutput struct {
	Running int `json:"running,omitempty"`
}

type toasterCountResponse struct {
	Success bool `json:"success"`
	Running int  `json:"running,omitempty"`
}

func (sess *Session) ToasterCount(input *ToasterCountInput) (*ToasterCountOutput, error) {
	resp := &toasterCountResponse{}

	if input.ID == "" {
		return nil, fmt.Errorf("you did not provide the ID of the Toaster")
	}

	apierr, err := sess.client.AuthedGet("/toaster/count/"+input.ID, resp)
	if err != nil {
		return nil, err
	}
	if apierr != nil {
		return nil, fmt.Errorf("APIERROR: status: %v; code: %v; message: %v", apierr.Status, apierr.Code, apierr.Message)
	}

	if !resp.Success {
		return nil, fmt.Errorf("The API returned a failure with a 200 HTTP status code which should not happen")
	}

	return &ToasterCountOutput{
		Running: resp.Running,
	}, nil
}

type ToasterStatsInput struct {
	ID string `json:"id,omitempty"`
}

type ToasterStatsOutput struct {
	Stats *models.ToasterStats `json:"stats,omitempty"`
}

type toasterStatsResponse struct {
	Success bool                 `json:"success"`
	Stats   *models.ToasterStats `json:"stats,omitempty"`
}

func (sess *Session) ToasterStats(input *ToasterStatsInput) (*ToasterStatsOutput, error) {
	resp := &toasterStatsResponse{}

	if input.ID == "" {
		return nil, fmt.Errorf("you did not provide the ID of the Toaster")
	}

	apierr, err := sess.client.AuthedGet("/toaster/stats/"+input.ID, resp)
	if err != nil {
		return nil, err
	}
	if apierr != nil {
		return nil, fmt.Errorf("APIERROR: status: %v; code: %v; message: %v", apierr.Status, apierr.Code, apierr.Message)
	}

	if !resp.Success {
		return nil, fmt.Errorf("The API returned a failure with a 200 HTTP status code which should not happen")
	}

	if resp.Stats == nil {
		return nil, fmt.Errorf("The request was successfull but the remote API returned an empty body")
	}

	return &ToasterStatsOutput{
		Stats: resp.Stats,
	}, nil
}

type GetToasterInput struct {
	ID string `json:"id,omitempty"`
}

type GetToasterOutput struct {
	Toaster *models.Toaster `json:"toaster,omitempty"`
}

type getToasterResponse struct {
	Success bool            `json:"success"`
	Toaster *models.Toaster `json:"toaster,omitempty"`
}

func (sess *Session) GetToaster(input *GetToasterInput) (*GetToasterOutput, error) {
	resp := &getToasterResponse{}

	if input.ID == "" {
		return nil, fmt.Errorf("you did not provide the ID of the Toaster to get")
	}

	apierr, err := sess.client.AuthedGet("/toaster/"+input.ID, resp)
	if err != nil {
		return nil, err
	}
	if apierr != nil {
		return nil, fmt.Errorf("APIERROR: status: %v; code: %v; message: %v", apierr.Status, apierr.Code, apierr.Message)
	}

	if !resp.Success {
		return nil, fmt.Errorf("The API returned a failure with a 200 HTTP status code which should not happen")
	}

	if resp.Toaster == nil {
		return nil, fmt.Errorf("The request was successfull but the remote API returned an empty body")
	}

	return &GetToasterOutput{
		Toaster: resp.Toaster,
	}, nil
}

type GetToasterFileInput struct {
	ID   string `json:"id,omitempty"`
	Path string `json:"path,omitempty"`
}

type GetToasterFileOutput struct {
	File io.ReadCloser
}

func (sess *Session) GetToasterFile(input *GetToasterFileInput) (*GetToasterFileOutput, error) {
	if input.ID == "" {
		return nil, fmt.Errorf("you did not provide the ID of the Toaster")
	}

	if len(input.Path) == 0 {
		return nil, fmt.Errorf("you did not provide the Path of the file to retrieve")
	}

	if input.Path[0] == '/' {
		if len(input.Path) == 1 {
			input.Path = ""
		} else {
			input.Path = input.Path[1:]
		}
	}

	if len(input.Path) == 0 {
		return nil, fmt.Errorf("you did not provide the Path of the file to retrieve")
	}

	file, apierr, err := sess.client.AuthedStreamedGet("/toaster/file/" + input.ID + "/" + input.Path)
	if err != nil {
		return nil, err
	}
	if apierr != nil {
		return nil, fmt.Errorf("APIERROR: status: %v; code: %v; message: %v", apierr.Status, apierr.Code, apierr.Message)
	}

	return &GetToasterFileOutput{
		File: file,
	}, nil
}

type ListToasterFilesInput struct {
	ID string `json:"id,omitempty"`
}

type ListToasterFilesOutput struct {
	Files []string `json:"files,omitempty"`
}

type listToasterFilesResponse struct {
	Success bool     `json:"success"`
	Files   []string `json:"files,omitempty"`
}

func (sess *Session) ListToasterFiles(input *ListToasterFilesInput) (*ListToasterFilesOutput, error) {
	resp := &listToasterFilesResponse{}

	if input.ID == "" {
		return nil, fmt.Errorf("you did not provide the ID of the Toaster to get")
	}

	apierr, err := sess.client.AuthedGet("/toaster/listfiles/"+input.ID, resp)
	if err != nil {
		return nil, err
	}
	if apierr != nil {
		return nil, fmt.Errorf("APIERROR: status: %v; code: %v; message: %v", apierr.Status, apierr.Code, apierr.Message)
	}

	if !resp.Success {
		return nil, fmt.Errorf("The API returned a failure with a 200 HTTP status code which should not happen")
	}

	return &ListToasterFilesOutput{
		Files: resp.Files,
	}, nil
}

type GetToasterLogsInput struct {
	ID    string `json:"id,omitempty"`
	ExeID string `json:"exe_id,omitempty"`
}

type GetToasterLogsOutput struct {
	Logs []byte `json:"logs,omitempty"`
}

type getToasterLogsResponse struct {
	Success bool   `json:"success"`
	Logs    []byte `json:"logs,omitempty"`
}

func (sess *Session) GetToasterLogs(input *GetToasterLogsInput) (*GetToasterLogsOutput, error) {
	resp := &getToasterLogsResponse{}

	if input.ID == "" {
		return nil, fmt.Errorf("you did not provide the ID of the Toaster to get")
	}

	apierr, err := sess.client.AuthedGet("/toaster/logs/"+input.ID+"/"+input.ExeID, resp)
	if err != nil {
		return nil, err
	}
	if apierr != nil {
		return nil, fmt.Errorf("APIERROR: status: %v; code: %v; message: %v", apierr.Status, apierr.Code, apierr.Message)
	}

	if !resp.Success {
		return nil, fmt.Errorf("The API returned a failure with a 200 HTTP status code which should not happen")
	}

	return &GetToasterLogsOutput{
		Logs: resp.Logs,
	}, nil
}

type ListToastersInput struct {
}

type ListToastersOutput struct {
	Toasters []models.Toaster `json:"toasters,omitempty"`
}

type listToastersResponse struct {
	Success  bool             `json:"success"`
	Toasters []models.Toaster `json:"toasters,omitempty"`
}

func (sess *Session) ListToasters(input *ListToastersInput) (*ListToastersOutput, error) {
	resp := &listToastersResponse{}

	apierr, err := sess.client.AuthedGet("/toaster/list", resp)
	if err != nil {
		return nil, err
	}
	if apierr != nil {
		return nil, fmt.Errorf("APIERROR: status: %v; code: %v; message: %v", apierr.Status, apierr.Code, apierr.Message)
	}

	if !resp.Success {
		return nil, fmt.Errorf("The API returned a failure with a 200 HTTP status code which should not happen")
	}

	return &ListToastersOutput{
		Toasters: resp.Toasters,
	}, nil
}

type DeleteToasterInput struct {
	IDs []string `json:"toaster_ids,omitempty"`
}

type DeleteToasterOutput struct {
}

type deleteToasterResponse struct {
	Success bool `json:"success"`
}

func (sess *Session) DeleteToaster(input *DeleteToasterInput) (*DeleteToasterOutput, error) {
	resp := &deleteToasterResponse{}

	apierr, err := sess.client.AuthedDelete("/toaster", input, resp)
	if err != nil {
		return nil, err
	}
	if apierr != nil {
		return nil, fmt.Errorf("APIERROR: status: %v; code: %v; message: %v", apierr.Status, apierr.Code, apierr.Message)
	}

	if !resp.Success {
		return nil, fmt.Errorf("The API returned a failure with a 200 HTTP status code which should not happen")
	}

	return &DeleteToasterOutput{}, nil
}

type CreateToasterInput struct {
	CryptoSecure bool `json:"cryptographically_secure,omitempty"`

	Codes     [][]byte `json:"codes,omitempty"`
	CodePaths []string `json:"code_paths,omitempty"`
	// OR
	CodeFolder string `json:"code_folder,omitempty"`
	// OR
	CodeStream chan *models.MultipartItem `json:"code_stream,omitempty"`
	// OR
	GitURL         string `json:"git_url,omitempty"`
	GitUsername    string `json:"git_username,omitempty"`
	GitAccessToken string `json:"git_access_token,omitempty"`
	GitPassword    string `json:"git_password,omitempty"`
	GitBranch      string `json:"git_branch,omitempty"`

	// Executed in the root directory of the codepaths
	BuildCmd []string `json:"build_command,omitempty"`
	ExeCmd   []string `json:"execution_command,omitempty"`
	Env      []string `json:"environment_variables,omitempty"`

	JoinableForSec       int `json:"joinable_for_seconds,omitempty"`
	MaxConcurrentJoiners int `json:"max_concurrent_joiners,omitempty"`
	TimeoutSec           int `json:"timeout_seconds,omitempty"`

	Name     string   `json:"name,omitempty"`
	Readme   string   `json:"readme,omitempty"`
	Keywords []string `json:"keywords,omitempty"`
}

type createToasterRequest struct {
	CryptoSecure bool `json:"cryptographically_secure,omitempty"`

	Codes          [][]byte `json:"codes,omitempty"`
	CodePaths      []string `json:"code_paths,omitempty"`
	GitURL         string   `json:"git_url,omitempty"`
	GitUsername    string   `json:"git_username,omitempty"`
	GitAccessToken string   `json:"git_access_token,omitempty"`
	GitPassword    string   `json:"git_password,omitempty"`
	GitBranch      string   `json:"git_branch,omitempty"`

	// Executed in the root directory where the codepaths have been put
	BuildCmd []string `json:"build_command,omitempty"`
	ExeCmd   []string `json:"execution_command,omitempty"`
	Env      []string `json:"environment_variables,omitempty"`

	JoinableForSec       int `json:"joinable_for_seconds,omitempty"`
	MaxConcurrentJoiners int `json:"max_concurrent_joiners,omitempty"`
	TimeoutSec           int `json:"timeout_seconds,omitempty"`

	Name     string   `json:"name,omitempty"`
	Readme   string   `json:"readme,omitempty"`
	Keywords []string `json:"keywords,omitempty"`
}

type CreateToasterOutput struct {
	Toaster   *models.Toaster `json:"toaster,omitempty"`
	Domain    string          `json:"domain,omitempty"`
	BuildLogs []byte          `json:"build_logs,omitempty"`
}

type createToasterResponse struct {
	Success   bool            `json:"success"`
	Toaster   *models.Toaster `json:"toaster,omitempty"`
	Domain    string          `json:"domain,omitempty"`
	BuildLogs []byte          `json:"build_logs,omitempty"`
}

func (sess *Session) CreateToaster(input *CreateToasterInput) (*CreateToasterOutput, error) {
	resp := &createToasterResponse{}
	req := &createToasterRequest{
		CryptoSecure:         input.CryptoSecure,
		BuildCmd:             input.BuildCmd,
		ExeCmd:               input.ExeCmd,
		Env:                  input.Env,
		JoinableForSec:       input.JoinableForSec,
		MaxConcurrentJoiners: input.MaxConcurrentJoiners,
		TimeoutSec:           input.TimeoutSec,
		Name:                 input.Name,
		Readme:               input.Readme,
		Keywords:             input.Keywords,
	}

	var err error
	var apierr *apiclient.Error
	switch {
	case len(input.CodePaths) > 0:
		req.Codes = input.Codes
		req.CodePaths = input.CodePaths
		apierr, err = sess.client.AuthedPost("/toaster", req, resp)
	case input.GitURL != "":
		req.GitURL = input.GitURL
		req.GitUsername = input.GitUsername
		req.GitAccessToken = input.GitAccessToken
		req.GitPassword = input.GitPassword
		req.GitBranch = input.GitBranch
		apierr, err = sess.client.AuthedPost("/toaster", req, resp)
	case input.CodeFolder != "":
		apierr, err = sess.client.AuthedMultipartFolderPost(input.CodeFolder, "/toaster", req, resp)
	case input.CodeStream != nil:
		apierr, err = sess.client.AuthedMultipartReadersPost(input.CodeStream, "/toaster", req, resp)
	default:
		apierr, err = sess.client.AuthedPost("/toaster", req, resp)
	}

	if err != nil {
		return nil, err
	}
	if apierr != nil {
		return nil, fmt.Errorf("APIERROR: status: %v; code: %v; message: %v", apierr.Status, apierr.Code, apierr.Message)
	}

	if !resp.Success {
		return nil, fmt.Errorf("The API returned a failure with a 200 HTTP status code which should not happen")
	}

	if resp.Toaster == nil {
		return nil, fmt.Errorf("The request was successfull but the remote API returned an empty body")
	}

	return &CreateToasterOutput{
		Toaster:   resp.Toaster,
		Domain:    resp.Domain,
		BuildLogs: resp.BuildLogs,
	}, nil
}

type UpdateToasterInput struct {
	ID string `json:"id,omitempty"`

	Codes     [][]byte `json:"codes,omitempty"`
	CodePaths []string `json:"code_paths,omitempty"`
	// OR
	CodeFolder string `json:"code_folder,omitempty"`
	// OR
	CodeStream chan *models.MultipartItem `json:"code_stream,omitempty"`
	// OR
	GitURL         string `json:"git_url,omitempty"`
	GitUsername    string `json:"git_username,omitempty"`
	GitAccessToken string `json:"git_access_token,omitempty"`
	GitPassword    string `json:"git_password,omitempty"`
	GitBranch      string `json:"git_branch,omitempty"`
	GitRefresh     bool   `json:"refresh_from_last_git,omitempty"`

	// Executed in the root directory of the codepaths
	BuildCmd []string `json:"build_command,omitempty"`
	ExeCmd   []string `json:"execution_command,omitempty"`
	Env      []string `json:"environment_variables,omitempty"`

	JoinableForSec       *int `json:"joinable_for_seconds,omitempty"`
	MaxConcurrentJoiners *int `json:"max_concurrent_joiners,omitempty"`
	TimeoutSec           *int `json:"timeout_seconds,omitempty"`

	Name     *string  `json:"name,omitempty"`
	Readme   *string  `json:"readme,omitempty"`
	Keywords []string `json:"keywords,omitempty"`
}

type updateToasterRequest struct {
	BuildCmd []string `json:"build_command,omitempty" bson:"build_command,omitempty"`
	ExeCmd   []string `json:"execution_command,omitempty" bson:"execution_command,omitempty"`
	Env      []string `json:"environment_variables,omitempty" bson:"environment_variables,omitempty"`

	JoinableForSec       *int `json:"joinable_for_seconds,omitempty" bson:"joinable_for_seconds,omitempty"`
	MaxConcurrentJoiners *int `json:"max_concurrent_joiners,omitempty" bson:"max_concurrent_joiners,omitempty"`
	TimeoutSec           *int `json:"timeout_seconds,omitempty" bson:"timeout_seconds,omitempty"`

	Name     *string  `json:"name,omitempty" bson:"name,omitempty"`
	Readme   *string  `json:"readme,omitempty" bson:"readme,omitempty"`
	Keywords []string `json:"keywords,omitempty" bson:"keywords,omitempty"`

	Codes          [][]byte `json:"codes,omitempty"`
	CodePaths      []string `json:"code_paths,omitempty"`
	GitURL         *string  `json:"git_url,omitempty"`
	GitUsername    *string  `json:"git_username,omitempty"`
	GitAccessToken *string  `json:"git_access_token,omitempty"`
	GitPassword    *string  `json:"git_password,omitempty"`
	GitBranch      *string  `json:"git_branch,omitempty"`
	GitRefresh     bool     `json:"refresh_from_last_git,omitempty"`
}

type UpdateToasterOutput struct {
	Toaster   *models.Toaster `json:"toaster,omitempty"`
	Domain    string          `json:"domain,omitempty"`
	BuildLogs []byte          `json:"build_logs,omitempty"`
}

type updateToasterResponse struct {
	Success   bool            `json:"success"`
	Toaster   *models.Toaster `json:"toaster,omitempty"`
	Domain    string          `json:"domain,omitempty"`
	BuildLogs []byte          `json:"build_logs,omitempty"`
}

func (sess *Session) UpdateToaster(input *UpdateToasterInput) (*UpdateToasterOutput, error) {
	if input.ID == "" {
		return nil, fmt.Errorf("you did not provide the ID of the Toaster to update")
	}

	resp := &updateToasterResponse{}
	req := &updateToasterRequest{
		BuildCmd:             input.BuildCmd,
		ExeCmd:               input.ExeCmd,
		Env:                  input.Env,
		JoinableForSec:       input.JoinableForSec,
		MaxConcurrentJoiners: input.MaxConcurrentJoiners,
		TimeoutSec:           input.TimeoutSec,
		Name:                 input.Name,
		Readme:               input.Readme,
		Keywords:             input.Keywords,
		GitRefresh:           input.GitRefresh,
	}

	var err error
	var apierr *apiclient.Error
	switch {
	case len(input.CodePaths) > 0:
		req.Codes = input.Codes
		req.CodePaths = input.CodePaths
		apierr, err = sess.client.AuthedPut("/toaster/"+input.ID, req, resp)
	case input.GitURL != "":
		req.GitURL = &input.GitURL
		req.GitUsername = &input.GitUsername
		req.GitAccessToken = &input.GitAccessToken
		req.GitPassword = &input.GitPassword
		req.GitBranch = &input.GitBranch
		apierr, err = sess.client.AuthedPut("/toaster/"+input.ID, req, resp)
	case input.CodeFolder != "":
		apierr, err = sess.client.AuthedMultipartFolderPut(input.CodeFolder, "/toaster/"+input.ID, req, resp)
	case input.CodeStream != nil:
		apierr, err = sess.client.AuthedMultipartReadersPut(input.CodeStream, "/toaster/"+input.ID, req, resp)
	default:
		apierr, err = sess.client.AuthedPut("/toaster/"+input.ID, req, resp)
	}

	if err != nil {
		return nil, err
	}
	if apierr != nil {
		return nil, fmt.Errorf("APIERROR: status: %v; code: %v; message: %v", apierr.Status, apierr.Code, apierr.Message)
	}

	if !resp.Success {
		return nil, fmt.Errorf("The API returned a failure with a 200 HTTP status code which should not happen")
	}

	if resp.Toaster == nil {
		return nil, fmt.Errorf("The request was successfull but the remote API returned an empty body")
	}

	return &UpdateToasterOutput{
		Toaster:   resp.Toaster,
		Domain:    resp.Domain,
		BuildLogs: resp.BuildLogs,
	}, nil
}
