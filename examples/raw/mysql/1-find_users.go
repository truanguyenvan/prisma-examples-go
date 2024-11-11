package mysql

import (
	"context"
	"fmt"
	"prisma-examples-go/prisma/db"
)

func FindUsers(ctx context.Context, client *db.PrismaClient) error {
	var users []db.UserModel
	err := client.Prisma.QueryRaw("SELECT * FROM `User`").Exec(ctx, &users)
	if err != nil {
		return err
	}

	fmt.Printf("users: %v", users)
	return nil
}
