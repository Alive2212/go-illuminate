package request

import (
	"github.com/labstack/echo"
	"strconv"
)

var QueryParamObject QueryParams

func GetQueryParams(ctx echo.Context) (QueryParams, error) {
	queryParams, err := GetPaginationFromRequest(ctx)
	if err != nil {
		return queryParams, err
	}
	queryParams, err = GetSortFromRequest(ctx)
	if err != nil {
		return queryParams, err
	}
	return queryParams, nil
}

func GetPaginationFromRequest(ctx echo.Context) (QueryParams, error) {
	page, err := strconv.Atoi(ctx.QueryParam("page"))
	if err != nil {
		QueryParamObject.Page = 1
	}
	QueryParamObject.Page = page
	limit, err := strconv.Atoi(ctx.QueryParam("per_page"))
	if err != nil {
		QueryParamObject.Limit = 1
	}
	QueryParamObject.Limit = limit
	return QueryParamObject, nil
}

func GetSortFromRequest(ctx echo.Context) (QueryParams, error) {
	sortBy := ctx.QueryParam("sort_by")
	QueryParamObject.SortBy = []string{sortBy}
	return QueryParamObject, nil
}
