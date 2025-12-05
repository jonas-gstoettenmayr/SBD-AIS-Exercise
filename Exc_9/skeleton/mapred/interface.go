package mapred

type MapReduceInterface interface {
	Run(input []string) map[string]int					// test split into lines -> get back words and how many words
	wordCountMapper(text string) []KeyValue				// array of single word and count 1
	wordCountReducer(key string, values []int) KeyValue // take single word and how many counts
}

// KeyValue represents a key-value pair
type KeyValue struct {
	Key   string
	Value int
}