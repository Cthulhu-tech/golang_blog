package query

import "github.com/Cthulhu-tech/golang_blog/internal/application/common"

type PostyQueryResult struct {
	Result *common.PostResult
}

type PostyQueryListResult struct {
	Result []*common.PostResult
}
