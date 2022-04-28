package adapter

type JsonHandle interface {
	handle()
}

type JsonHandler struct{}

func (j JsonHandler) handle() {
	return
}

type XMLHandle interface {
	handle()
}

