package ws

type Cmd struct {
	Op   string        `json:"op"`
	Args []interface{} `json:"args"`
}
