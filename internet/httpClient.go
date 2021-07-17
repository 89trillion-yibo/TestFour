package internet

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func HttpClient(uuid string,gifcode string) []byte {
	params := url.Values{"id": {uuid},"gifcode":{gifcode}}
	resp, _ := http.PostForm("http://localhost:8080/reward", params)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return body
}
