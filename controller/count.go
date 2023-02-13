package controller

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	coreResponse "github.com/datsukan/datsukan-blog-comment-core/response"
	"github.com/datsukan/datsukan-blog-comment-core/usecase"
	"github.com/datsukan/datsukan-blog-comment-count/request"
	"github.com/datsukan/datsukan-blog-comment-count/response"
)

// Count は、コメント数を取得するControllerの処理。
func Count(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	req := request.GetRequest(r.QueryStringParameters)
	if err := req.Validate(); err != nil {
		return coreResponse.ResponseBadRequestError(err)
	}

	c, err := usecase.Count(req.ArticleID)
	if err != nil {
		return coreResponse.ResponseInternalServerError(err)
	}

	res := response.Response{
		ArticleID: req.ArticleID,
		Count:     int(c),
	}

	j, err := json.Marshal(res)
	if err != nil {
		return coreResponse.ResponseInternalServerError(err)
	}
	js := string(j)

	return coreResponse.ResponseSuccess(js)
}
