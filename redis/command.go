package redis

import (
	"context"
)

// ExistsAndSet judge this id exists or not
// Using `setnx` to discern, and add an eventIDPrefix
func ExistsAndSet(id string) (bool, error) {
	initClient()

	doSet, err := client.SetNX(context.Background(), eventIDPrefix+id, "", 0).Result()
	return !doSet, err
}
