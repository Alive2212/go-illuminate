package response

import (
	"github.com/biezhi/gorm-paginator/QueryParamsObject"
	"github.com/labstack/echo"
	"math"
)

func Response(ctx echo.Context, response ResponseModel) error {
	return ctx.JSON(response.ResponseCode, map[string]interface{}{
		"response": response,
	})
}

func SuccessfulResponse(ctx echo.Context, response SuccessfulResponseModel) error {
	return ctx.JSON(200, successfulResponse{
		true,
		response.Data,
		GetTrans(response.MessageKey),
	})
}

func GetTrans(messageKey string) string {
	// TODO implement here for translate
	return messageKey
}

// response bulk
func BulkResponse(paginator *QueryParamsObject.Paginator, ctx echo.Context) error {
	pagesCount := int(math.Ceil(float64(paginator.TotalRecord) / float64(paginator.Limit)))
	hasNext := bool(paginator.Page != pagesCount)
	result := ctx.JSON(200, SuccessfulBulk{
		Result: paginator.Records,
		Pagination: Pagination{
			Page:       paginator.Page,
			PerPage:    paginator.Limit,
			TotalCount: paginator.TotalRecord,
			PagesCount: pagesCount,
			HasNext:    hasNext,
		},
	})
	return result
}