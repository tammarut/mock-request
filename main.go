package main

import (
	"fmt"
	"net/http"
)

//HTTPClient is interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

//ClientMock is struct
type ClientMock struct {
}

func main() {
	//=>Mock
	mock := &ClientMock{}

	//=>Call func
	err := MyClient(mock)
	if err != nil { //=>Handle error
		fmt.Println(err.Error())
	}
}

//MyClient for set-up request
func MyClient(client HTTPClient) error {
	request, err := http.NewRequest("POST", "https://www.reallycoolurl.com", nil)
	if err != nil {
		return err
	}

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	fmt.Println(response.Status)
	fmt.Println(response.StatusCode)
	fmt.Println("Successful request.")
	return nil
}

//Do is Do
func (c *ClientMock) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
	}, nil
}
