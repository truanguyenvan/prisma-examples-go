package main

import (
	"context"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
	"prisma-examples-go/examples"
	"prisma-examples-go/examples/raw/mongo/commandRaw"
	"prisma-examples-go/examples/raw/mongo/findRaw"
	"prisma-examples-go/examples/raw/mysql"

	"prisma-examples-go/prisma/db"
)

var (
	dbType     string
	DSN        string
	schemaGen  bool
	exampleNum string
	err        error

	SchemaPath = "prisma/schema.prisma"
	MongoDSN   = "mongodb://prisma:pw@localhost:27016/prisma?authSource=admin&retryWrites=true"
	MySQLDSN   = "mysql://root:pw@localhost:3307/testing"
)

func initPrismaClient(DNS string) *db.PrismaClient {
	return db.NewClient(db.WithDatasourceURL(DNS))
}

// runExampleCmd represents the run-example command
var runExampleCmd = &cobra.Command{
	Use:   "run-example",
	Short: "Run Prisma example with specified database type",
	Long:  `Run Prisma example for MongoDB, MySQL, or other supported databases.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Example logic to run Prisma commands
		runPrismaExample(dbType)
	},
}

func init() {
	// Define the --db flag with a default value of "mongodb"
	runExampleCmd.Flags().StringVarP(&dbType, "db", "d", "mongodb", "Database type to use (mongodb, mysql, etc.)")
	runExampleCmd.Flags().StringVarP(&DSN, "dsn", "s", "", "Database DSN to use")
	runExampleCmd.Flags().BoolVarP(&schemaGen, "schema-gen", "g", false, "Generate Prisma schema")
	runExampleCmd.Flags().StringVarP(&exampleNum, "example", "e", "no example", "Example number to run")
}

func runPrismaExample(dbType string) {
	// //go:generate go run github.com/steebchen/prisma-client-go generate --schema prisma/mongo_schema.prisma
	switch dbType {
	case "mongodb":
		SchemaPath = "prisma/mongo_schema.prisma"
		if DSN == "" {
			DSN = MongoDSN
		}
	case "mysql":
		SchemaPath = "prisma/mysql_schema.prisma"
		if DSN == "" {
			DSN = MySQLDSN
		}
	default:
		log.Fatalf("Unsupported database type: %s", dbType)
	}

	log.Printf("Running Prisma example for database: %s, DSN: %s, SchemaPath: %s\n", dbType, DSN, SchemaPath)

	if schemaGen {
		log.Println("Generating Prisma client")
		if res, err := exec.Command("go", "run", "github.com/steebchen/prisma-client-go", "generate", "--schema", SchemaPath).CombinedOutput(); err != nil {
			log.Fatalf("Error generating Prisma client:%s -  %s", string(res), err)
		}
	}

	client := initPrismaClient(DSN)
	if err := client.Connect(); err != nil {
		log.Fatalf("Error connecting to Prisma client: %s", err)
	}

	defer func() {
		if err := client.Disconnect(); err != nil {
			log.Fatalf("Error disconnecting from Prisma client: %s", err)
		}
	}()

	ctx := context.Background()
	log.Printf("Running Prisma example: %s\n", exampleNum)
	switch exampleNum {
	// Example 1: Create a new user
	case "1":
		err = examples.CreateUser(ctx, client)
	// Example 2: Find a user by Email
	case "2":
		err = examples.FindUsers(ctx, client)
	// Example 3: Update a user's email
	case "3":
		err = examples.UpdateUser(ctx, client)
	// Example 4: Delete a user
	case "4":
		err = examples.DeleteUser(ctx, client)
	// Example 5: Create a new post
	case "5":
		err = examples.CreateUserPosts(ctx, client)
	// Example 6: Find posts by user
	case "6":
		err = examples.FindPosts(ctx, client)

	// Raw queries
	// mysql
	case "raw-mysql-1":
		err = mysql.FindUsers(ctx, client)
	case "raw-mysql-2":
		err = mysql.CreateUser(ctx, client)
	// mongodb
	case "raw-mongo-command-1":
		err = commandRaw.CreateUser(ctx, client)
	case "raw-mongo-command-2":
		err = commandRaw.FindUsers(ctx, client)
	case "raw-mongo-find-1":
		err = findRaw.FindUsers(ctx, client)
	default:
		log.Fatalf("Unsupported example number: %s", exampleNum)
	}

	if err != nil {
		log.Fatalf("Error running Prisma example: %s", err)
	}
}

func main() {
	if err := runExampleCmd.Execute(); err != nil {
		log.Fatalf("Error running Prisma example: %s", err)
	}
}
