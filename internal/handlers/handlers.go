package handlers

type HttpHandlers struct {
	UserHandler UserHandler
}

func NewHttpHandlers(userHandler UserHandler) *HttpHandlers {
	return &HttpHandlers{
		UserHandler: userHandler,
	}
}
