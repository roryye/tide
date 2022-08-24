package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"

	"golang.org/x/sync/errgroup"
)

var SalesBaseURL = `https://interview.moreless.io/questions/async_workers/sales/`

// QuerySalesData 获取销售数据
func QuerySalesData(ctx context.Context) (int, error) {
	urls := []string{
		SalesBaseURL + "a",
		SalesBaseURL + "b",
		SalesBaseURL + "c",
		SalesBaseURL + "d",
	}

	var sum int
	mu := sync.Mutex{}
	g, ctx := errgroup.WithContext(ctx)
	for _, url := range urls {
		g.Go(func() error {
			count, err := getCount(ctx, url)
			if err != nil {
				return err
			}
			mu.Lock()
			sum += count
			mu.Unlock()
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return -1, err
	}
	return sum, nil
}

func getCount(ctx context.Context, url string) (int, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return -1, err
	}

	type data struct {
		Count int    `json:"count"`
		Code  int    `json:"code"`
		Msg   string `json:"msg"`
	}
	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return -1, err
	}
	d := data{}
	if err = json.Unmarshal(b, &d); err != nil {
		return -1, err
	}
	fmt.Println("data", d)
	if d.Code != 0 {
		return -1, errors.New(d.Msg)
	}
	return d.Count, nil
}
