package kodi

//type PlaylistId int

type Playlist struct {
	k *Kodi
}

func (p *Playlist) doRPC(rpcName string, params interface{}, response interface{}) error {
	req := &request{
		Jsonrpc: "2.0",
		Method:  "Playlist." + rpcName,
		Id:      1,
		Params:  params,
	}
	return p.k.doRPC(req, response)
}

// Playlist.Add

type PlaylistAddParams struct {
	PlaylistId PlaylistId `json:"playlistid"`
	// TODO: http://kodi.wiki/view/JSON-RPC_API/v6#Playlist.Add claims that item
	// below is just a single value, but from Kodi's responses it looks like it
	// accepts a list too.
	Item PlaylistItem `json:"item"`
}

type PlaylistAddResult string

type PlaylistAddResponse struct {
	ResponseBase
	Result *PlaylistAddResult `json:"result,omitempty"`
}

func (p *Playlist) Add(params *PlaylistAddParams) (*PlaylistAddResponse, error) {
	response := &PlaylistAddResponse{}
	err := p.doRPC("Add", params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Playlist.Clear

type PlaylistClearParams struct {
	PlaylistId PlaylistId `json:"playlistid"`
}

type PlaylistClearResult string

type PlaylistClearResponse struct {
	ResponseBase
	Result PlaylistClearResult `json:"result,omitempty"`
}

func (p *Playlist) Clear(params *PlaylistClearParams) (*PlaylistClearResponse, error) {
	response := &PlaylistClearResponse{}
	err := p.doRPC("Clear", params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Playlist.GetItems

type PlaylistGetItemsParams struct {
	PlaylistId PlaylistId      `json:"playlistid"`
	Properties []ListFieldsAll `json:"properties,omitempty"`
	Limits     *ListLimits     `json:"limits,omitempty"`
	Sort       *ListSort       `json:"sort,omitempty"`
}

type PlaylistGetItemsResult struct {
	Limits ListLimitsReturned `json:"limits"`
	Items  []*ListItemAll     `json:"items"`
}

type PlaylistGetItemsResponse struct {
	ResponseBase
	Result *PlaylistGetItemsResult `json:"result,omitempty"`
}

func (p *Playlist) GetItems(params *PlaylistGetItemsParams) (*PlaylistGetItemsResponse, error) {
	response := &PlaylistGetItemsResponse{}
	err := p.doRPC("GetItems", params, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Playlist.GetPlaylists

type PlaylistGetPlaylistsResult struct {
	PlaylistId   PlaylistId   `json:"playlistid"`
	PlaylistType PlaylistType `json:"playlisttype"`
}

type PlaylistGetPlaylistsResponse struct {
	ResponseBase
	Result []*PlaylistGetPlaylistsResult `json:"result,omitempty"`
}

func (p *Playlist) GetPlaylists() (*PlaylistGetPlaylistsResponse, error) {
	response := &PlaylistGetPlaylistsResponse{}
	err := p.doRPC("GetPlaylists", nil, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Playlist.GetProperties

type PlaylistGetPropertiesParams struct {
	PlaylistId PlaylistId             `json:"playlistid"`
	Properties []PlaylistPropertyName `json:"properties,omitempty"`
}

type PlaylistGetPropertiesResult PlaylistPropertyValue

type PlaylistGetPropertiesResponse struct {
	ResponseBase
	Result *PlaylistGetPropertiesResult `json:"result,omitempty"`
}

func (p *Playlist) GetProperties() (*PlaylistGetPropertiesResponse, error) {
	response := &PlaylistGetPropertiesResponse{}
	err := p.doRPC("GetProperties", nil, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Playlist.Insert

type PlaylistInsertParams struct {
	PlaylistId PlaylistId   `json:"playlistid"`
	Position   int          `json:"position"`
	Item       PlaylistItem `json:"item"`
}

type PlaylistInsertResult string

type PlaylistInsertResponse struct {
	ResponseBase
	Result *PlaylistInsertResponse `json:"result,omitempty"`
}

func (p *Playlist) Insert() (*PlaylistInsertResponse, error) {
	response := &PlaylistInsertResponse{}
	err := p.doRPC("Insert", nil, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Playlist.Remove

type PlaylistRemoveParams struct {
	PlaylistId PlaylistId   `json:"playlistid"`
	Position   int          `json:"position"`
	Item       PlaylistItem `json:"item"`
}

type PlaylistRemoveResult string

type PlaylistRemoveResponse struct {
	ResponseBase
	Result *PlaylistRemoveResult `json:"result,omitempty"`
}

func (p *Playlist) Remove() (*PlaylistRemoveResponse, error) {
	response := &PlaylistRemoveResponse{}
	err := p.doRPC("Remove", nil, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Playlist.Swap

type PlaylistSwapParams struct {
	PlaylistId PlaylistId `json:"playlistid"`
	Position1  int        `json:"position1"`
	Position2  int        `json:"position2"`
}

type PlaylistSwapResult string

type PlaylistSwapResponse struct {
	ResponseBase
	Result *PlaylistSwapResult `json:"result,omitempty"`
}

func (p *Playlist) Swap() (*PlaylistSwapResponse, error) {
	response := &PlaylistSwapResponse{}
	err := p.doRPC("Swap", nil, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
