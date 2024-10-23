package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Keys struct {
	GUID       uuid.UUID `gorm:"primaryKey" json:"guid"`
	BusinessId string    `json:"business_id"`
	PrivateKey string    `json:"private_key"`
	PublicKey  string    `json:"public_key"`
	Timestamp  uint64
}

type KeysView interface {
	QueryKeysByBusId(string, uint64, uint64) ([]Keys, error)
}

type KeysDB interface {
	KeysView

	StoreKeys([]Keys, uint64) error
}

type addressesDB struct {
	gorm *gorm.DB
}

func NewKeysDB(db *gorm.DB) KeysDB {
	return &addressesDB{gorm: db}
}

func (db *addressesDB) StoreKeys(keyList []Keys, keyLengthe uint64) error {
	result := db.gorm.CreateInBatches(&keyList, int(keyLengthe))
	return result.Error
}

func (db *addressesDB) QueryKeysByBusId(busId string, page, pageSize uint64) ([]Keys, error) {
	panic("implement me")
}
