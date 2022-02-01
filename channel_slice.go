// Generated by: gen
// TypeWriter: slice
// Directive: +gen on Channel

package main

// ChannelSlice is a slice of type Channel. Use it where you would use []Channel.
type ChannelSlice []Channel

// GroupByString groups elements into a map keyed by string. See: http://clipperhouse.github.io/gen/#GroupBy
func (rcv ChannelSlice) GroupByString(fn func(Channel) string) map[string]ChannelSlice {
	result := make(map[string]ChannelSlice)
	for _, v := range rcv {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}