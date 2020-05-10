package webservices

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// WebService contain data to call a service
type WebService struct {
	endpoint string
	headers  map[string]string
}

// New return un pointeur sur un struc WebService
func New(endpoint string, mapHeaders map[string]string) *WebService {

	webservice := &WebService{
		endpoint: endpoint,
		headers:  mapHeaders,
	}
	return webservice
}

// CallWebService appel le webservice passé en paramètre
func (w *WebService) CallWebService() error {

	req, _ := http.NewRequest("POST", w.endpoint, nil)

	for key, value := range w.headers {
		req.Header.Add(key, value)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	return nil
}
