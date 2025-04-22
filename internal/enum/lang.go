package enum

type Lang string

const (
	JP Lang = "japanese"
)

func (l Lang) OpenAI() string {
	switch l {
	case JP:
		return "ja"
	default:
		return ""
	}
}
