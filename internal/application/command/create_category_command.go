package command

import (
	"time"

	"github.com/Cthulhu-tech/golang_blog/internal/application/common"
)

type CreateCategoryCommand struct {
	ID        string
	Name      string
	CreatedAt time.Time
}

type CreateCategoryCommandResult struct {
	Result *common.CategoryResult
}
