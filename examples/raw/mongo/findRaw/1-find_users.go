package findRaw

import (
	"context"
	"fmt"
	"prisma-examples-go/prisma/db"
)

func FindUsers(ctx context.Context, client *db.PrismaClient) error {
	var users []db.UserModel
	client.User.FindRaw()
	fmt.Printf("users: %v", users)
	return nil
}
