package examples

import (
	"context"
	"encoding/json"
	"fmt"
	"prisma-examples-go/prisma/db"
)

func CreateUserPosts(ctx context.Context, client *db.PrismaClient) error {
	// create user
	createdUser, err := client.User.CreateOne(
		db.User.Email.Set("admin@gmail"),
		db.User.FirstName.Set("Trua Nguyen"),
		db.User.LastName.Set("admin"),
	).Exec(ctx)
	if err != nil {
		return err
	}

	// create posts
	createdPost1, err := client.Post.CreateOne(
		db.Post.Title.Set("Post 1"),
		db.Post.Body.Set("Body 1"),
		db.Post.Author.Link(
			db.User.ID.Equals(createdUser.ID),
		),
	).Exec(ctx)
	if err != nil {
		return err
	}

	createdPost2, err := client.Post.CreateOne(
		db.Post.Title.Set("Post 2"),
		db.Post.Body.Set("Body 2"),
		db.Post.Author.Link(
			db.User.ID.Equals(createdUser.ID),
		),
	).Exec(ctx)
	if err != nil {
		return err
	}

	// print result
	result, _ := json.MarshalIndent(map[string]interface{}{
		"user": createdUser,
		"posts": []interface{}{
			createdPost1,
			createdPost2,
		},
	}, "", "  ")
	fmt.Printf("created user and posts: %s\n", result)
	return nil
}
