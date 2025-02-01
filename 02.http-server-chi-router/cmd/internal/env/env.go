package env

// READ MORE: https://direnv.net
// Install: brew install direnv
// Add to zshrc: eval "$(direnv hook zsh)"
// Approve content: direnv allow
import (
	"os"
	"strconv"
)

func GetString(key string, fallback string) string {
	val, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	return val
}

func GetInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}

	return valAsInt
}
