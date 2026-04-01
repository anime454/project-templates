package in

import (
	"context"
)

type VehicleService interface {
	CheckIn(ctx context.Context, licensePlate string) error
	CheckOut(ctx context.Context, licensePlate string) error
}
