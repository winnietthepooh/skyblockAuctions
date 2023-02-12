package auctions

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"sync"
)

func DoAll() *AuctionContainer {
	init, _ := http.NewRequest(http.MethodGet, "https://api.hypixel.net/skyblock/auctions", nil)
	response, err := http.DefaultClient.Do(init)

	if err != nil {
		return nil
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil
	}
	var fullAuctions FullAuctionData
	err = json.Unmarshal(body, &fullAuctions)

	container := AuctionContainer{
		Data:  fullAuctions.Auctions,
		Mutex: sync.Mutex{},
	}

	var wg sync.WaitGroup
	for i := 1; i < fullAuctions.TotalPages; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			auc := request(i)
			container.Mutex.Lock()
			container.Data = append(container.Data, auc...)
			container.Mutex.Unlock()
		}()
	}

	wg.Wait()

	return &container
}

func request(page int) []Auctions {
	init, _ := http.NewRequest(http.MethodGet, "https://api.hypixel.net/skyblock/auctions", nil)
	init.URL.RawQuery = "page=" + strconv.Itoa(page)
	// fmt.Println(init.URL)
	response, err := http.DefaultClient.Do(init)
	if err != nil {
		return nil
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil
	}
	var fullAuctions FullAuctionData
	err = json.Unmarshal(body, &fullAuctions)
	return fullAuctions.Auctions
}
