package query

import "github.com/Cthulhu-tech/golang_blog/internal/application/common"

type CategoryQueryResult struct {
	Result *common.CategoryResult
}

type CategoryQueryListResult struct {
	Result []*common.CategoryResult
}
