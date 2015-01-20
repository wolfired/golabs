package cli

type Command struct {
	Key     string
	Options map[string]Option
}

type Option struct {
	Key string
}
