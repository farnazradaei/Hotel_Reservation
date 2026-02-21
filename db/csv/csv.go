package csv

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"sync"
)

type User struct {
	ID    int
	Email string
	Hash  string
}

type CSVStore struct {
	path string
	mu   sync.Mutex
}

func NewCsv(path string) *CSVStore {
	return &CSVStore{
		path: path,
	}
}

func (c *CSVStore) UserExsist(email string) (bool, error) {
	f, err := os.Open("data/userdata.csv")
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return false, err
	}

	for i, row := range records {
		if i == 0 {
			continue
		}

		if row[1] == email {
			return true, nil
		}
	}

	return false, nil

}

func (c *CSVStore) NextID() (int, error) {
	f, err := os.Open("data/userdata.csv")
	if err != nil {
		return 1, nil
	}
	defer f.Close()

	reader := csv.NewReader(f)
	rows, _ := reader.ReadAll()

	if len(rows) <= 1 {
		return 1, nil
	}

	last := rows[len(rows)-1][0]
	id, err := strconv.Atoi(last)
	if err != nil {
		return 1, nil
	}
	return id + 1, nil
}

var mu sync.Mutex

func (c *CSVStore) UserSave(u User) error {
	mu.Lock()
	defer mu.Unlock()

	f, err := os.OpenFile("data/userdata.csv",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := csv.NewWriter(f)

	fileInfo, _ := f.Stat()
	if fileInfo.Size() == 0 {
		writer.Write([]string{"ID", "Email", "Hash"})
	}

	err = writer.Write([]string{
		strconv.Itoa(u.ID),
		u.Email,
		u.Hash,
	})

	if err != nil {
		return err
	}

	writer.Flush()
	return writer.Error()
}

func GetUserByEmail(email string) (User, error) {
	f, err := os.Open("data/userdata.csv")
	if err != nil {
		return User{}, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return User{}, err
	}

	for i, row := range records {
		if i == 0 {
			continue
		}
		if row[1] == email {
			id, _ := strconv.Atoi(row[0])
			return User{
				ID:    id,
				Email: row[1],
				Hash:  row[2],
			}, nil
		}
	}

	return User{}, errors.New("user not found")
}
