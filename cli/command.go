package cli





type Command struct {
	key string
	options map[string]Option
}

type Option struct {
	key string
}
