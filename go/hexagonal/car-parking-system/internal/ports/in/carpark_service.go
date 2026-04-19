package in

import (
	"context"

	"github.com/anime454/project-templates/go/hexagonal/car-parking-system/internal/domain"
)

type CarParkService interface {
	CheckIn(ctx context.Context, licensePlate string) (domain.Vehicle, error)
	CheckOut(ctx context.Context, licensePlate string) error
}
