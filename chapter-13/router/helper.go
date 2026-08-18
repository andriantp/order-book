package router

import "hash/fnv"

func hashMarket(market string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(market))
	return h.Sum32()
}
