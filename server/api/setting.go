package api

type DrbdEnableSetting struct {
	Enabledrbd bool   `json:"enabledrbd"`
	State      State  `json:"state"`
	Version    string `json:"version"`
}
