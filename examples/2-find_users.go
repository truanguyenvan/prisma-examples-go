package examples

import (
	"context"
	"fmt"
	"prisma-examples-go/prisma/db"
)

func FindUsers(ctx context.Context, client *db.PrismaClient) error {
	// find all users
	users, err := client.User.FindMany(db.User.Email.Equals("truanv@gmail")).Exec(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("users: %v\n", users)
	return nil
}
