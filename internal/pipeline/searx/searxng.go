package searx

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	domain "smartsearch/internal/entities"
)

func CallSearx(query string) (*domain.Searx, error) {
	q := fmt.Sprintf(
		"http://localhost:8080/search?q=%s&format=json&count=%d",
		query,
		10, // max URLs you want
	)
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
