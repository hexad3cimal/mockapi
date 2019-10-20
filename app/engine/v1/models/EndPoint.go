package models

type EndPoint struct {
	Id     string `json:"id"`
	ApiId     string `json:"apiId"`
	URL string `json:"url"`
	Method string `json:"method"`
	RequestBody string `json:"requestBody"`
	StatusCode int      `json:"statusCode"`
	Headers string `json:"headers"`
	ResponseBody  string `json:responseBody"`
}