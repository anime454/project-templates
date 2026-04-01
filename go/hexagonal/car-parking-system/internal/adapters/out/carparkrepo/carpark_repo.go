package vehiclerepo

import (
	"context"

	"github.com/anime454/project-templates/go/hexagonal/car-parking-system/internal/adapters/out/db"
	"github.com/anime454/project-templates/go/hexagonal/car-parking-system/internal/domain"
)

type CarparkRepository struct {
	db *db.Adaptor
}

func NewCarparkRepository(db *db.Adaptor) *CarparkRepository {
	return &CarparkRepository{
		db: db,
	}
}

func (r *CarparkRepository) CheckExistingVehicle(ctx context.Context, vehicle domain.Vehicle) (domain.Vehicle, error) {
	var existingVehicle domain.Vehicle
	err := r.db.WithContext(ctx).Find(&existingVehicle, map[string]any{
		"license_plate": vehicle.LicensePlate,
	})
	return existingVehicle, err
}
