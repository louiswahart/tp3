package agt

type AgentID int64
type Alternative int

type AgentI interface {
	Equal(ag Agent) bool
	DeepEqual(ag Agent) bool
	Clone() AgentI
	String() string
	Prefers(a Alternative, b Alternative)
	Start()
}

type Agent struct {
	ID    AgentID
	Name  string
	Prefs []Alternative
}
