package aggregateRaw

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"prisma-examples-go/prisma/db"
)

func FindUsers(ctx context.Context, client *db.PrismaClient) error {
	users, err := client.User.AggregateRaw(bson.A{
		bson.M{"$match": bson.M{"email": bson.M{"$eq": "truanv@gmail"}, "lastName": bson.M{"$eq": "Van"}}},
		bson.M{"$project": bson.M{"email": 1, "lastName": 1, "_id": 0}},
	}).Exec(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("users: %v", users)
	return nil
}
