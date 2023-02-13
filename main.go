package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/datsukan/datsukan-blog-comment-core/usecase"
	"github.com/datsukan/datsukan-blog-comment-count/controller"
	"github.com/joho/godotenv"
)

func main() {
	t := flag.Bool("local", false, "ローカル実行か否か")
	articleID := flag.String("article-id", "", "ローカル実行用の記事ID")
	flag.Parse()

	isLocal, err := isLocal(t, articleID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if isLocal {
		fmt.Println("local")
		localController(articleID)
		return
	}

	fmt.Println("production")
	lambda.Start(controller.Count)
}

// isLocal はローカル環境の実行であるかを判定する。
func isLocal(t *bool, articleID *string) (bool, error) {
	if !*t {
		return false, nil
	}

	if *articleID == "" {
		fmt.Println("no exec")
		return false, fmt.Errorf("ローカル実行だが記事ID指定が無いので処理不可能")
	}

	return true, nil
}

// localController はローカル環境での実行処理を行う。
func localController(articleID *string) {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("環境変数を読み込み出来ませんでした: %v\n", err)
		return
	}

	c, err := usecase.Count(*articleID)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("count: %d\n\n", c)
}
