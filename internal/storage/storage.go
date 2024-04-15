package storage

import (
	"fmt"

	"github.com/FedorKowarnow/storage/internal/file"
	"github.com/google/uuid"
)

type Storage struct {
	iles map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		iles: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(filename string, blop []byte) (*file.File, error) {
	newFile, err := file.NewFile(filename, blop)
	if err != nil {
		return nil, err
	}
	s.iles[newFile.ID] = newFile
	return newFile, nil
}

func (s *Storage) GetByID(fileID uuid.UUID) (*file.File, error) {
	foundFile, ok := s.iles[fileID]
	if !ok {
		return nil, fmt.Errorf("File %v not found", fileID)
	}
	return foundFile, nil
}
