package examples

import (
	"context"
	"fmt"
	"prisma-examples-go/prisma/db"
)

func DeleteUser(ctx context.Context, client *db.PrismaClient) error {
	// delete user
	deletedUser, err := client.User.FindMany(
		db.User.Email.Equals("truanv@gmail"),
	).Delete().Exec(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("deleted user: %v\n", deletedUser)
	return nil
}
