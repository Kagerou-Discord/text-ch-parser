package main

import orderedmap "github.com/wk8/go-ordered-map"

// GroupByString groups elements into a map keyed by string. See: http://clipperhouse.github.io/gen/#GroupBy
func (rcv ChannelSlice) GroupByStringExt(fn func(Channel) string) orderedmap.OrderedMap {
	mm := orderedmap.New()
	for _, v := range rcv {
		key := fn(v)
		a, b := mm.Get(key)
		if !b {
			var ss ChannelSlice
			mm.Set(key, append(ss, v))
		} else {
			mm.Set(key, append(a.(ChannelSlice), v))
		}
	}
	return *mm
}
