package config

type Config struct {
	KeyBindings []KeyBinding
}
type KeyBinding struct {
	Key    string
	Mods   string
	Action string
	Mode   string
}
