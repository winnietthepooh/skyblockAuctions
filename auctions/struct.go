package auctions

import (
	"sync"
)

type AuctionContainer struct {
	Data  []Auctions
	Mutex sync.Mutex
}

type FullAuctionData struct {
	Success       bool       `json:"success"`
	Page          int        `json:"page"`
	TotalPages    int        `json:"totalPages"`
	TotalAuctions int        `json:"totalAuctions"`
	LastUpdated   int64      `json:"lastUpdated"`
	Auctions      []Auctions `json:"auctions"`
}

type Auctions struct {
	Uuid             string        `json:"uuid"`
	Auctioneer       string        `json:"auctioneer"`
	ProfileId        string        `json:"profile_id"`
	Coop             []string      `json:"coop"`
	Start            int64         `json:"start"`
	End              int64         `json:"end"`
	ItemName         string        `json:"item_name"`
	ItemLore         string        `json:"item_lore"`
	Extra            string        `json:"extra"`
	Category         string        `json:"category"`
	Tier             string        `json:"tier"`
	StartingBid      int           `json:"starting_bid"`
	ItemBytes        string        `json:"item_bytes"`
	Claimed          bool          `json:"claimed"`
	ClaimedBidders   []interface{} `json:"claimed_bidders"`
	HighestBidAmount int           `json:"highest_bid_amount"`
	LastUpdated      int64         `json:"last_updated"`
	Bin              bool          `json:"bin"`
	Bids             []bids        `json:"bids"`
	ItemUuid         string        `json:"item_uuid,omitempty"`
}

type bids struct {
	AuctionId string `json:"auction_id"`
	Bidder    string `json:"bidder"`
	ProfileId string `json:"profile_id"`
	Amount    int    `json:"amount"`
	Timestamp int64  `json:"timestamp"`
}
