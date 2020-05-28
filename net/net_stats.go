package net

type BasePktStat struct {
	Bits uint64 //bits statis
	Pkts uint64 //packet statis
}

type NetEventStat struct {
	Valid BasePktStat
	Error BasePktStat
	Drop  BasePktStat
	Total BasePktStat
}
