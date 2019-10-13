package environment

import "github.com/joho/godotenv"

func Load(filePath string) (err error) {
	return godotenv.Load(filePath)
}
