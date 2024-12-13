package index

import (
	"context"

	db_index "github.com/whosonfirst/go-whosonfirst-database/app/sql/index"
)

func Run(ctx context.Context) error {
	return db_index.Run(ctx)
}
