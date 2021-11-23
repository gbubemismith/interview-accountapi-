package accounts

type Mock struct {
	Method      string
	Url         string
	RequestBody string

	Error        error
	ResponseBody string
	StatusCode   int
}
