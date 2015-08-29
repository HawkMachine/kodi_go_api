package kodi

// http://kodi.wiki/view/JSON-RPC_API/v6#Player

type Player struct {
	k *Kodi
}

func (p *Player) doRPC(rpcName string, params interface{}, response interface{}) error {
	req := &request{
		Jsonrpc: "2.0",
		Method:  "Player." + rpcName,
		Id:      "1",
		Params:  params,
	}
	return p.k.doRPC(req, response)
}

// Player.GetActivePlayers

type PlayerGetActiveResultElement struct {
	Type PlayerType `json:"type"`
	PlayerId
}

type PlayerGetActivePlayersResponse struct {
	ResponseBase
	Result []*PlayerGetActiveResultElement
}

// Player.GetItem

type PlayerGetItemParams struct {
	PlayerId   PlayerId         `json:"playerid"`
	Properties []*ListFieldsAll `json:"properties"`
}

type PlayerGetItemResult struct {
	Item ListItemAll `json:"item"`
}
