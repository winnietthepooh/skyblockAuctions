package memory

import "skyblockAuctions/auctions"

var auctionsDB = make(map[string]auctions.Auctions)

func AddAuction(auction auctions.Auctions) {
	auctionsDB[auction.Uuid] = auction
}

func RemoveAuction(uuid string) {
	delete(auctionsDB, uuid)
}

func GetAuction(uuid string) (auctions.Auctions, bool) {
	auction, ok := auctionsDB[uuid]
	return auction, ok
}
