package helpers

type IStrinigHelper interface {
	GetSlugString(str string, delimiter *string) string
}
