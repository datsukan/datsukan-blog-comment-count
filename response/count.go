package response

// Response は、コメント数の取得結果の構造体。
type Response struct {
	ArticleID string `json:"article_id"`
	Count     int    `json:"count"`
}
