package searx

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	domain "smartsearch/internal/domain"
)

func CallSearx(query string) (*domain.Searx, error) {
	q := fmt.Sprintf("http://localhost:8080/search?q=%s&format=json", query)
	resp, err := http.Get(q)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	var urlData domain.Searx
	err = json.Unmarshal(body, &urlData)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return &urlData, nil
}
