package examples

import (
	"context"
	"encoding/json"
	"fmt"
	"prisma-examples-go/prisma/db"
)

func FindPosts(ctx context.Context, client *db.PrismaClient) error {
	// find a single user
	user, err := client.User.FindUnique(
		db.User.Email.Equals("admin@gmail"),
	).With(db.User.Posts.Fetch().Take(3)).Exec(ctx)
	if err != nil {
		return err
	}

	fmt.Println("--------------user: ", user)

	result, _ := json.MarshalIndent(user, "", "  ")
	fmt.Printf("post: %s\n", result)

	return nil
}
