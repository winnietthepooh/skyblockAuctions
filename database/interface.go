package database

import "skyblockAuctions/auctions"

type Database interface {
	Add() error
	Get() (auctions.Auctions, error)
}
