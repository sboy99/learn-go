package ports

// ------------------------------INTERFACES---------------------------------- //

type ICryptoHelperPort interface {
	HashStr(str string) (string, error)
	CompareHash(str string, hash string) bool
}
