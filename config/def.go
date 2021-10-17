package config

func GetDefaultConfig() Config {
	result := Config{
		KeyBindings: getDefaultKeyBindings(),
	}
	return result
}

func getDefaultKeyBindings() []KeyBinding {
	var result []KeyBinding
	result = append(result, KeyBinding{
		Key:    "j",
		Action: "moveDown",
	})
	result = append(result, KeyBinding{
		Key:    "k",
		Action: "moveUp",
	})
	result = append(result, KeyBinding{
		Key:    "h",
		Action: "moveLeft",
	})
	result = append(result, KeyBinding{
		Key:    "l",
		Action: "moveRight",
	})

	return result
}
