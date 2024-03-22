package lsp

type InitializeRequest struct {
	Request
	Params InitializeRequestParams `json:"params"`
}

type InitializeRequestParams struct {
	ClientInfo *ClientInfo `json:"clientInfo"`
	// TODO: there are more params to implement
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
