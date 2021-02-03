package core

type MachineInfos struct {
	MachineInfos []Infos `json:"infos"`
}

type Infos struct {
	port    int    `json:"port"`
	process string `json:"process"`
}
