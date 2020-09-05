package util

type RealTarget string

func NewTarget(target string) RealTarget {
	switch target {
	case "block":
		return "block"
	case "mute":
		return "mute"
	case "follower":
		return "follower"
	case "followee":
		return "followee"
	default:
		return "unknown"
	}
}

func (t RealTarget) GetFileName() string {
	return t.String() + "list.txt"
}

func (t RealTarget) String() string {
	return string(t)
}
