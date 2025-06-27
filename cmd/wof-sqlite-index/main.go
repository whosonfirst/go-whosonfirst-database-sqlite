package main

import (
	"context"
	"log"

	_ "github.com/whosonfirst/go-whosonfirst-database-sqlite"
	_ "github.com/whosonfirst/go-whosonfirst-iterate-sql/v3"
	
	"github.com/whosonfirst/go-whosonfirst-database/app/sql/tables/index"
)

func main() {

	ctx := context.Background()
	err := index.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to index, %v", err)
	}
}
