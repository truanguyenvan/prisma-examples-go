package examples

import (
	"context"
	"encoding/json"
	"fmt"
	"prisma-examples-go/prisma/db"
)

func CreateUser(ctx context.Context, client *db.PrismaClient) error {
	// create use
	createdUser, err := client.User.CreateOne(
		db.User.Email.Set("truanv@gmail"),
		db.User.FirstName.Set("Trua Nguyen"),
		db.User.LastName.Set("Van"),
	).Exec(ctx)
	if err != nil {
		return err
	}

	result, _ := json.MarshalIndent(createdUser, "", "  ")
	fmt.Printf("created user: %s\n", result)
	return nil
}
