package fault

type APIResponseMessage int

const (
	NotFoundResponse APIResponseMessage = iota
)

var messageMap = map[APIResponseMessage]string{
	NotFoundResponse: "Entity was not found",
}

func (arm APIResponseMessage) String() string {
	return messageMap[arm]
}
