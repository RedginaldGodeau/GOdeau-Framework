package Methods

type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	HEAD   Method = "HEAD"
	DELETE Method = "DELETE"
	PUT    Method = "PUT"
	PATCH  Method = "PATCH"
)

func List(method ...Method) []Method {
	return method
}
