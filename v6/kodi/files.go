package kodi

// Implementation of Files RPCs.
type Files struct {
	k *Kodi
}

func (f *Files) doRPC(rpcName string, params interface{}, response interface{}) error {
	req := &request{
		Jsonrpc: "2.0",
		Method:  "Files." + rpcName,
		Id:      1,
		Params:  params,
	}
	return f.k.doRPC(req, response)
}

// 5.4.2 Files.GetDirectory
// Get the directories and files in the given directory.

type FilesGetDirectoryParams struct {
	Directory string      `json:"directory"`
	Media     *FilesMedia `json:"media,omitempty"`

	Properties []ListFieldsFiles `json:"properties,omitempty"`
	Sort       *ListSort         `json:"sort,omitempty"`
}

type FilesGetDirectoryResult struct {
	Limits ListLimitsReturned `json:"limits"`
	Files  []ListItemFile     `json:"files"`
}

type FilesGetDirectoryResponse struct {
	ResponseBase
	Result *FilesGetDirectoryResult
}

func (f *Files) GetDirectory(params *FilesGetDirectoryParams) (*FilesGetDirectoryResponse, error) {
	response := &FilesGetDirectoryResponse{}
	err := f.doRPC("GetDirectory", params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// 5.4.4 Files.GetSources

type FilesGetSourcesParams struct {
	Media  FilesMedia  `json:"media,omitempty"`
	Limits *ListLimits `json:"limits,omitempty"`
	Sort   *ListSort   `json:"sort,omitempty"`
}

type FilesGetSourcesResult struct {
	Limits  ListLimitsReturned `json:"limits"`
	Sources []ListItemsSource  `json:"sources"`
}

type FilesGetSourcesResponse struct {
	ResponseBase
	Result *FilesGetSourcesResult `json:"result,omitempty"`
}

func (f *Files) GetSources(params *FilesGetSourcesParams) (*FilesGetSourcesResponse, error) {
	response := &FilesGetSourcesResponse{}
	err := f.doRPC("GetSources", params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
