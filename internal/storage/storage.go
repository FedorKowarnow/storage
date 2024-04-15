package storage

import "github.com/FedorKowarnow/storage/internal/file"

type Storage struct{}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Upload(filename string, blop []byte) (*file.File, error) {
	return file.NewFile(filename, blop)
}
