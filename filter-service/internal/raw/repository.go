package raw

import "filter-service/pkg/db"

type RawRepository struct {
	Database *db.Db
}

func NewRawRepository(database *db.Db) *RawRepository {
	return &RawRepository{
		Database: database,
	}
}

func (r *RawRepository) Save(rawData *RawData) (*RawData, error) {
	resutl := r.Database.DB.Create(rawData)
	if resutl.Error != nil {
		return nil, resutl.Error
	}
	return rawData, nil
}
