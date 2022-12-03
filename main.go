package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"time"

	pkgcf "github.com/datsukan/contentful-good-decrement-lambda/pkg/contentful"
	pkgga "github.com/datsukan/contentful-good-decrement-lambda/pkg/goodattr"
	pkgrp "github.com/datsukan/contentful-good-decrement-lambda/pkg/resparse"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	contentful "github.com/contentful-labs/contentful-go"
)

type Response struct {
	GoodCount int `json:"goodCount"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

var (
	cma   *contentful.Client
	space *contentful.Space
)

var headers = map[string]string{
	"Access-Control-Allow-Origin":  "*",
	"Access-Control-Allow-Methods": "PUT",
	"Access-Control-Allow-Headers": "Content-Type",
}

func main() {
	t := flag.Bool("local", false, "ローカル実行か否か")
	ID := flag.String("id", "", "ローカル実行用の記事ID")
	flag.Parse()

	isLocal, err := isLocal(t, ID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if isLocal {
		fmt.Println("local")
		localController(ID)
		return
	}

	fmt.Println("production")
	lambda.Start(controller)
}

func isLocal(t *bool, ID *string) (bool, error) {
	if !*t {
		return false, nil
	}

	if *ID == "" {
		fmt.Println("no exec")
		return false, fmt.Errorf("ローカル実行だがID指定が無いので処理不可能")
	}

	return true, nil
}

func localController(ID *string) {
	js, err := useCase(*ID)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(js)
}

func controller(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	articleID := request.PathParameters["article_id"]
	if articleID == "" {
		err := fmt.Errorf("article_id is empty")
		return responseBadRequestError(err)
	}

	js, err := useCase(articleID)
	if err != nil {
		return responseInternalServerError(err)
	}

	return responseSuccess(js)
}

func useCase(articleID string) (string, error) {
	// Contentful SDK のクライアントインスタンスを生成する
	var err error
	cma, space, err = pkgcf.NewContentfulSDK()
	if err != nil {
		return "", err
	}

	var g int
	g, err = decrement(articleID)
	if err != nil {
		for i := 0; i < 4; i++ {
			time.Sleep(time.Millisecond * 500)

			g, err = decrement(articleID)
			if err == nil {
				break
			}
		}
		if err != nil {
			return "", err
		}
	}

	r := Response{
		GoodCount: g,
	}
	jb, err := json.Marshal(r)
	if err != nil {
		return "", err
	}

	return string(jb), nil
}

func responseBadRequestError(rerr error) (events.APIGatewayProxyResponse, error) {
	b := ErrorResponse{
		Error:   "bad request",
		Message: rerr.Error(),
	}
	jb, err := json.Marshal(b)
	if err != nil {
		r := events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers:    headers,
			Body:       err.Error(),
		}
		return r, nil
	}
	body := string(jb)

	r := events.APIGatewayProxyResponse{
		StatusCode: 400,
		Headers:    headers,
		Body:       body,
	}
	return r, nil
}

func responseInternalServerError(rerr error) (events.APIGatewayProxyResponse, error) {
	b := ErrorResponse{
		Error:   "internal server error",
		Message: rerr.Error(),
	}
	jb, err := json.Marshal(b)
	if err != nil {
		r := events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers:    headers,
			Body:       err.Error(),
		}
		return r, nil
	}
	body := string(jb)

	r := events.APIGatewayProxyResponse{
		StatusCode: 500,
		Headers:    headers,
		Body:       body,
	}
	return r, nil
}

func responseSuccess(body string) (events.APIGatewayProxyResponse, error) {
	r := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    headers,
		Body:       body,
	}
	return r, nil
}

func decrement(articleID string) (int, error) {
	article, err := article(articleID)
	if err != nil {
		return 0, err
	}

	g, err := goods(article)
	if err != nil {
		return 0, err
	}

	// いいね数をデクリメント
	var rg int
	if g <= 0 {
		rg = 0
	} else {
		rg = g - 1
	}

	// Contentfulから取得した記事情報に、ローカルで管理されている記事の本文を反映する
	fg, err := pkgrp.IntToField(rg)
	if err != nil {
		return 0, err
	}
	article.Fields["goods"] = fg

	// ローカルで管理されている記事の本文を、Contentfulへ反映する
	if err := cma.Entries.Upsert(space.Sys.ID, article); err != nil {
		return 0, err
	}

	return rg, nil
}

func article(articleID string) (*contentful.Entry, error) {
	article, err := cma.Entries.Get(space.Sys.ID, articleID)
	if err != nil {
		return nil, err
	}

	// Contentfulから記事情報が取得できない場合、処理を終了する
	if article == nil {
		return nil, fmt.Errorf("article not found")
	}

	return article, nil
}

func goods(article *contentful.Entry) (int, error) {
	// いいね数を取得する
	g, err := pkgga.GoodsAttr(article)
	if err != nil {
		return 0, err
	}

	return g, nil
}
