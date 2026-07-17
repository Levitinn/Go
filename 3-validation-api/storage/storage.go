package storage

import (
	"encoding/json"
	"os"
)

type Storage struct {
	path string // ./storage/tokens.json
}

func NewStorage(path string) *Storage {
	return &Storage{path: path}
}
func (s *Storage) Save(hash string, email string) error {
	data, err := os.ReadFile(s.path)
	tokens := make(map[string]string)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else if len(data) > 0 {
		err = json.Unmarshal(data, &tokens)
		if err != nil {
			return err
		}
	}
	tokens[hash] = email
	bytes, err := json.Marshal(tokens)
	if err != nil {
		return err
	}
	return os.WriteFile(s.path, bytes, 0644)
}
func (s *Storage) Find(hash string) (string, bool) {
	data, err := os.ReadFile(s.path)
	tokens := make(map[string]string)
	if err != nil {
		return "", false
	}
	if len(data) > 0 {
		err = json.Unmarshal(data, &tokens)
		if err != nil {
			return "", false
		}
	}
	email, ok := tokens[hash]
	return email, ok
}

func (s *Storage) Delete(hash string) error {
	data, err := os.ReadFile(s.path)
	tokens := make(map[string]string)
	if err != nil {
		return err
	}
	if len(data) > 0 {
		err = json.Unmarshal(data, &tokens)
		if err != nil {
			return err
		}
	}
	delete(tokens, hash)
	bytes, err := json.Marshal(tokens)
	if err != nil {
		return err
	}
	return os.WriteFile(s.path, bytes, 0644)
}
