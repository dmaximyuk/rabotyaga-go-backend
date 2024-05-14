package types

type EventType = string
type ErrorMessage = string

const (
	RequestPing     EventType = "ping"
	RequestStartApp EventType = "start_app"
	RequestBalance  EventType = "balance"
	RequestUser     EventType = "user"
)

const (
	ResponsePong     EventType = "pong"
	ResponseStartApp EventType = "start_app"
	ResponseBalance  EventType = "balance"
	ResponseUser     EventType = "user"
	ResponseError    EventType = "error"
)

const (
	//ErrorMessageNil          ErrorMessage = "data is nil"
	ErrorMessageParseData ErrorMessage = "data is nil"
	ErrorMessageMsgLength ErrorMessage = "data is nil"
	//ErrorMessageUidUndefined ErrorMessage = "the user ID is not defined or a conversion error has occurred"
	ErrorMessageMissingEvent ErrorMessage = "event is missing"
	ErrorMessageMissingData  ErrorMessage = "data is missing"
)
