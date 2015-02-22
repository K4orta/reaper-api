package stats

import (
	"os"
)

func FormatUrl(url string) string {
	return url + "?locale=en_US&apiKey=" + os.Getenv("BNET_API_KEY")
}
