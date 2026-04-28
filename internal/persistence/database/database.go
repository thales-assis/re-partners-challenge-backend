package database

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/re-partners-challenge-backend/internal/domain/entity"
	"github.com/re-partners-challenge-backend/internal/infra/log"
)

func ProvideDatabase(
	logger *log.ZapLogger,
) *FakeDatabase {

	file, err := os.Open("./internal/persistence/database/fake-database.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	var fakeDatabaseTotalItems FakeDatabaseTotalItems
	if err := json.NewDecoder(file).Decode(&fakeDatabaseTotalItems); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}

	result := make(map[uint32]entity.Pack)

	for _, pkg := range fakeDatabaseTotalItems.Records {

		newPack := entity.Pack{
			ID:        pkg.ID,
			CreatedAt: pkg.CreatedAt,
			Size:      pkg.Size,
		}

		result[newPack.ID] = newPack
	}

	return &FakeDatabase{
		CountIDs: 3,
		Records:  result,
	}
}
