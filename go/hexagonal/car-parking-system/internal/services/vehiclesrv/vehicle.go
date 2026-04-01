package vehiclesrv

import (
	"context"

	"github.com/anime454/project-templates/go/hexagonal/car-parking-system/internal/ports/out"
)

type VehicleService struct {
	vehicleRepo out.VehicleRepository
}

func NewVehicleService(vehicleRepo out.VehicleRepository) *VehicleService {
	return &VehicleService{
		vehicleRepo: vehicleRepo,
	}
}

func (s *VehicleService) CheckIn(ctx context.Context, licensePlate string) error {
	return nil
}
