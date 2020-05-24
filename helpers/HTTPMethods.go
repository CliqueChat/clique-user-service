package helpers

type HTTPMethod int

const (
	GET HTTPMethod = iota
	POST
	DELETE
)

var httpMethods = [...]string{
	"GET",
	"POST",
	"DELETE",
}

func (h HTTPMethod) String() string { return httpMethods[h] }
