package tools

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func ResponseMap(url string) (interface{}, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	req.Header.Add("x-rapidapi-key", "d31b9e03797eed17b1bd020146af292d")
	req.Header.Add("x-rapidapi-host", "v1.baseball.api-sports.io")

	res, httpErr := client.Do(req)
	if httpErr != nil {
		fmt.Print(httpErr)
		return nil, httpErr
	}

	defer res.Body.Close()

	bodyBytes, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		fmt.Print(readErr)
		return nil, readErr
	}

	var data struct {
		Data interface{}
	}

	unmarshalErr := json.Unmarshal(bodyBytes, &data.Data)
	if unmarshalErr != nil {
		fmt.Print(unmarshalErr)
		return nil, unmarshalErr
	}

	if data.Data == nil {
		data.Data = "Is Empty"
	}

	return data.Data, nil
}

func Response(url string, c *gin.Context) {

	var data struct {
		Data interface{}
	}

	data.Data = requests[url]
	if data.Data != nil {
		c.JSON(http.StatusOK, data.Data)
	} else {
		client := &http.Client{}

		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Error"})
			fmt.Print(err)
			return
		}

		req.Header.Add("x-rapidapi-key", "d31b9e03797eed17b1bd020146af292d")
		req.Header.Add("x-rapidapi-host", "v1.baseball.api-sports.io")

		res, httpErr := client.Do(req)
		if httpErr != nil {
			c.JSON(http.StatusBadRequest, httpErr)
			return
		}

		defer res.Body.Close()

		bodyBytes, readErr := io.ReadAll(res.Body)
		if readErr != nil {
			c.JSON(http.StatusBadRequest, readErr)
			return
		}

		unmarshalErr := json.Unmarshal(bodyBytes, &data.Data)
		if unmarshalErr != nil {
			c.JSON(http.StatusBadRequest, unmarshalErr)
			return
		}

		if data.Data == nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Is Empty"})
		} else {
			fmt.Print(data.Data)
			c.JSON(http.StatusOK, data.Data)
		}
	}

	requests[url] = data.Data

}
