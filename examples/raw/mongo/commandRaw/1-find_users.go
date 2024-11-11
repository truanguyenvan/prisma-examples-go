package commandRaw

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"prisma-examples-go/prisma/db"
)

func FindUsers(ctx context.Context, client *db.PrismaClient) error {
	var users []db.UserModel
	err := client.Prisma.RunCommandRaw(bson.M{"find": "User"}).Exec(ctx, &users)
	if err != nil {
		return err
	}

	fmt.Printf("users: %v", users)
	return nil
}
