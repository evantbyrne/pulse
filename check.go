package pulse

import (
	"fmt"
	"net/http"
)

func Check(url string) error {
	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return fmt.Errorf("Unexpected status code %d", response.StatusCode)
	}

	return nil
}
