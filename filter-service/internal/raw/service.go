package raw

import (
	"filter-service/commom"
	"log"
)

type RawService struct {
	rawRepository   *RawRepository
	kafkaRepository *KafkaRepository
}

func NewRawService(rawRepo *RawRepository, kafkaRepo *KafkaRepository) *RawService {
	return &RawService{
		rawRepository:   rawRepo,
		kafkaRepository: kafkaRepo,
	}
}

func (s *RawService) SaveToBD(raw, valid string) error {
	rawData := RawData{
		Data:  raw,
		Valid: valid,
	}
	if rawData.Valid == "true" {
		go func() {
			errk := s.kafkaRepository.Send([]byte(rawData.Data))
			if errk != nil {
				log.Println(commom.ErrSendToKafka, errk)
			}
		}()

	}
	_, err := s.rawRepository.Save(&rawData)
	if err != nil {
		log.Println(commom.ErrSaveToBDRAw, err)
		return err
	}
	return nil
}
