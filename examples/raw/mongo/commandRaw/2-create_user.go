package commandRaw

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"prisma-examples-go/prisma/db"
)

func CreateUser(ctx context.Context, client *db.PrismaClient) error {
	var insertedUser interface{}
	err := client.Prisma.RunCommandRaw(bson.M{"insert": "User", "documents": bson.M{"email": "truanv@gmail", "first_name": "Trua Nguyen ", "last_name": "Van"}}).Exec(ctx, &insertedUser)
	if err != nil {
		return err
	}

	fmt.Printf("insertedUser: %v", insertedUser)
	return nil
}
