package examples

import (
	"context"
	"fmt"
	"prisma-examples-go/prisma/db"
)

func UpdateUser(ctx context.Context, client *db.PrismaClient) error {
	// update user
	updatedUser, err := client.User.FindMany(
		db.User.Email.Equals("truanv@gmail"),
	).Update(
		db.User.FirstName.Set("Trua"),
	).Exec(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("updated user: %v\n", updatedUser)
	return nil
}
