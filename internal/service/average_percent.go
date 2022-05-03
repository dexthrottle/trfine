package service

// import (
// 	"context"

// 	"github.com/dexthrottle/trfine/internal/model"
// 	"github.com/dexthrottle/trfine/pkg/logging"
// 	"gorm.io/gorm"
// )

// type AveragePercentService interface {
// 	InsertAveragePercent(ctx context.Context, a model.AveragePercent) (*model.AveragePercent, error)
// }

// type averagePercentService struct {
// 	ctx        context.Context
// 	connection *gorm.DB
// 	log        logging.Logger
// }

// func NewAveragePercentService(ctx context.Context, db *gorm.DB, log logging.Logger) AveragePercentService {
// 	return &averagePercentService{
// 		ctx:        ctx,
// 		connection: db,
// 		log:        log,
// 	}
// }

// func (db *averagePercentService) InsertAveragePercent(ctx context.Context, a model.AveragePercent) (*model.AveragePercent, error) {
// 	tx := db.connection.WithContext(ctx)
// 	res := tx.Save(&a)
// 	db.log.Error(res.Error)
// 	if res.Error != nil {
// 		return nil, res.Error
// 	}
// 	return &a, nil
// }
