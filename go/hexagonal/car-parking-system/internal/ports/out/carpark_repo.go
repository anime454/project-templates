package out

import (
	"context"

	"github.com/anime454/project-templates/go/hexagonal/car-parking-system/internal/domain"
)

type CarparkRepository interface {
	CheckExistingVehicle(ctx context.Context, licensePlate string) (domain.Vehicle, error)
}
