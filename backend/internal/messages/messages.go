package messages

type MSG struct {
	Message string `json:"message"`
}

type ERR struct {
	Error string `json:"error"`
}

func Message(msg string) MSG {
	return MSG{msg}
}

func Error(err error) ERR {
	return ERR{err.Error()}
}
