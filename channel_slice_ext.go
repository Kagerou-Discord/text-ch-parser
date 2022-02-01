package main

import orderedmap "github.com/wk8/go-ordered-map"

// GroupByStringExt groups elements into a map keyed by string. Change return type of `gen#GroupByString` to `orderedmap.OrderedMap`
func (rcv ChannelSlice) GroupByStringExt(fn func(Channel) string) orderedmap.OrderedMap {
	om := orderedmap.New()
	for _, ch := range rcv {
		key := fn(ch)
		chSlice, exists := om.Get(key)

		if exists {
			om.Set(key, append(chSlice.(ChannelSlice), ch))
		} else {
			var slice ChannelSlice
			om.Set(key, append(slice, ch))
		}
	}
	return *om
}
