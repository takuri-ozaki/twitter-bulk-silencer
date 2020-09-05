package util

type Mode string

func NewMode(mode string) Mode {
	if mode == "block" {
		return "block"
	}
	if mode == "mute" {
		return "mute"
	}
	return "unknown"
}

func (m Mode) IsBlockMode() bool {
	return m == "block"
}

func (m Mode) String() string {
	return string(m)
}
