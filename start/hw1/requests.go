package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func hedgedcurlStart(ctx context.Context, a *Args, results chan<- string) {
	for _, url := range a.Urls {
		go func(u string) {
			req, err := http.NewRequestWithContext(ctx, "GET", u, nil)
			if err != nil {
				return
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				return
			}
			defer resp.Body.Close()

			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				return
			}

			var headerBuilder strings.Builder
			for name, values := range resp.Header {
				if len(values) > 0 {
					headerBuilder.WriteString(fmt.Sprintf("%s: %s\n", name, values[0]))
				}
			}

			resStr := fmt.Sprintf(
				"URL: %s\nHTTP/1.1 %s\n\n--- HEADERS ---\n%s\n--- BODY (Length: %d) ---\n%s\n",
				u, resp.Status, headerBuilder.String(), len(bodyBytes), string(bodyBytes),
			)

			select {
			case results <- resStr:
			case <-ctx.Done():
				return
			}
		}(url)
	}
}
