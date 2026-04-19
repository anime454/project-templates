package carParksrv

import (
	"context"

	"github.com/anime454/project-templates/go/hexagonal/car-parking-system/internal/adapters/out/db"
	"github.com/anime454/project-templates/go/hexagonal/car-parking-system/internal/domain"
	"github.com/anime454/project-templates/go/hexagonal/car-parking-system/internal/ports/out"
	"github.com/anime454/project-templates/go/hexagonal/car-parking-system/internal/services"
)

type CarParkService struct {
	carParkRepo out.CarparkRepository
}

func NewCarParkService(carParkRepo out.CarparkRepository) *CarParkService {
	return &CarParkService{
		carParkRepo: carParkRepo,
	}
}

func (s *CarParkService) CheckIn(ctx context.Context, licensePlate string) (domain.Vehicle, error) {
	existingVehicle, err := s.carParkRepo.CheckExistingVehicle(ctx, licensePlate)
	if err != nil && err != db.ErrRecordNotFound {
		return domain.Vehicle{}, services.ErrDatabaseQueryError
	}

	if existingVehicle.LicensePlate != "" {
		return existingVehicle, services.ErrVehicleAlreadyCheckedIn
	}

	return domain.Vehicle{}, nil
}
