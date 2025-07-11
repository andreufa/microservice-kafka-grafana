package raw

import (
	"filter-service/internal/kafkapkg"
	"filter-service/pkg/db"
)

type RawRepository struct {
	Database *db.Db
}
type KafkaRepository struct {
	Producer *kafkapkg.Producer
}

func NewRawRepository(database *db.Db) *RawRepository {
	return &RawRepository{
		Database: database,
	}
}
func NewKafkaReposotiry(p *kafkapkg.Producer) *KafkaRepository {
	return &KafkaRepository{
		Producer: p,
	}
}

func (r *RawRepository) Save(rawData *RawData) (*RawData, error) {
	resutl := r.Database.DB.Create(rawData)
	if resutl.Error != nil {
		return nil, resutl.Error
	}
	return rawData, nil
}

func (k *KafkaRepository) Send(data []byte) error {
	err := k.Producer.Produce(data)
	return err
}
