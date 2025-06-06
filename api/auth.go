import (
	"os"
	"github.com/joho/godotenv"
)

func retrieveAuthToken() (string, error) {
	envFile, err := godotenv.Read(".env")
	if err != nil {
		return err
	}

	token := os.Getenv("TOKEN")

	return token, nil
}