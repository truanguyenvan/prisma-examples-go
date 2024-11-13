package findRaw

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"prisma-examples-go/prisma/db"
)

func FindUsers(ctx context.Context, client *db.PrismaClient) error {
	users, err := client.User.FindRaw(bson.M{"email": bson.M{"$eq": "truanv@gmail"}, "lastName": bson.M{"$eq": "Van"}}).Exec(ctx)
	if err != nil {
		return err
	}

	fmt.Printf("users: %v", users)
	return nil
}
