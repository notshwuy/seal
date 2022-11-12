package shared

import (
	"os"

	"github.com/joho/godotenv"
)

func ReadEnvironmentFile(Path string) map[string]string {
	env, error := godotenv.Read(Path + "/.env")

	if error != nil {
		// defaults to empty environment
		return map[string]string{}
	}

	return env
}

func WorkingDirectory() string {
	current_path, error := os.Getwd()

	if error != nil {
		panic(error)
	}

	return current_path
}
