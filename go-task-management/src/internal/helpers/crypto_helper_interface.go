package helpers

type ICryptoHelper interface {
	GetId() int64
	GetHash(str string) string
	GetRandStr(size int) string
	GetRandInt(min int, max int) int
}
