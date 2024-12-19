package parsing

type Farm struct {
	NbrAnts string
	Start   string
	End     string
	Rooms   map[string][]string
	Tunnels map[string][]string
}

func NewGraph() *Farm {
	return &Farm{
		Rooms:   make(map[string][]string),
		Tunnels: make(map[string][]string),
	}
}
