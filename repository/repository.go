package repository

type Repository interface {
	Read() ([]byte, error)
	Write(data []byte)
}
