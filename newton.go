package newton

import (
	"encoding/json"
	"net/http"
	"strings"
)

//Data wraps the response from the newton API.
//The first 2 properties are provided by the user. The last property is provided by the API.
type Data struct {
	Operation, Expression, Result string
}

func buildURL(operation, expression string) string {
	var url strings.Builder
	url.WriteString("https://newton.now.sh/")
	url.WriteString(operation)
	url.WriteString("/")
	url.WriteString(expression)
	return url.String()
}

var client = &http.Client{}

func getResult(url string, target interface{}) error {
	r, err := client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

//Calc sends the expression and operation off to Newton then returns the resulting Data struct
func Calc(requestData Data) Data {
	result := Data{}
	err := getResult(buildURL(requestData.Operation, requestData.Expression), &result)
	if err != nil {
		panic(err)
	}
	return result
}
