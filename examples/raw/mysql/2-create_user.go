package mysql

import (
	"context"
	"fmt"
	"prisma-examples-go/prisma/db"
)

func CreateUser(ctx context.Context, client *db.PrismaClient) error {
	createdUser, err := client.Prisma.ExecuteRaw("INSERT INTO `User` (email, first_name, last_name) VALUES ('truanv@gmail', 'Trua Nguyen ', 'Van')").Exec(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("created user: %v", createdUser)
	return nil
}
