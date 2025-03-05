package query

import "github.com/Cthulhu-tech/golang_blog/internal/application/common"

type TagQueryResult struct {
	Result *common.TagResult
}

type TagQueryListResult struct {
	Result []*common.TagResult
}
