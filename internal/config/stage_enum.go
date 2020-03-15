package config

type Stage int

const (
	StageDev Stage = iota
	StageSIT
)

func (s Stage) String() string {
	return [...]string{"dev", "sit"}[s]
}
