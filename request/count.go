package request

import "fmt"

// GetRequest は、コメント取得のリクエストの構造体。
type Request struct {
	ArticleID string
}

func GetRequest(r map[string]string) *Request {
	return &Request{ArticleID: r["article_id"]}
}

func (r *Request) Validate() error {
	if r.ArticleID == "" {
		return fmt.Errorf("article_id is empty")
	}

	return nil
}
