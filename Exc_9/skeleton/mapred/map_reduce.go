package mapred
import (
	"sync"
	"strings"
)

type MapReduce struct {
}

// implement mapreduce

// a-z no other characters are "letters" for the test
func is_letter(r rune) rune {
		if (r >= 'a' && r <= 'z') || r == ' ' { // rune is i32 wrapper
			return r 
		}
		return -1 // deletes the character
}

// output list of each individual word (many copies)
func (MapReduce) wordCountMapper(text string) []KeyValue {
	keyvalues := []KeyValue{}
	text = strings.Map(is_letter, strings.ToLower(text)) // only a-z and whitespaces are allowed
	for _, word := range strings.Fields(text) { // like split but doesn't care about whitespace amount
		keyvalues = append(keyvalues, KeyValue {Key: word, Value: 1} )
	}
	return keyvalues
}


// sums up the values we gete for a key
func (MapReduce) wordCountReducer(key string, values []int) KeyValue {
	if len(values) < 1 {
		return KeyValue{Key: key, Value: 0}
	}
	result := KeyValue{Key: key, Value: values[0]}
	for _, val := range values[1:]{
		result.Value += val
	}
	return result
}

// take a list and make a map of the amount of occurances with the word
// used for reducing the output of wordCountMapper from (john, 1), (john, 1) to (john, [1,1,])
func shuffle(list []KeyValue) map[string][]int {
	if len(list) < 1{return make(map[string][]int)}

	wordsOccurences := make(map[string][]int)
	for _, item := range list {
		if _, exists := wordsOccurences[item.Key]; !exists{
			wordsOccurences[item.Key] = []int{item.Value}
		} else {
			wordsOccurences[item.Key] = append(wordsOccurences[item.Key], item.Value)
		}
	}
	return wordsOccurences
}

func (mr MapReduce) Run(input []string) map[string]int {
	
	num_goroutines := len(input) 
	
	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup
	var wg3 sync.WaitGroup
	lineWords := make(chan []KeyValue)
	wordMaps := make(chan map[string][]int)

	// start in reverse order for maximum efficiency
	// i.e. starting to listen before we send
	// 3 step process
	// all 3 steps are done for every line
	// the merging of the results is synchronisly after
	// not most beautifull way to code this, but is more parralizable this way

	// 3. get lists of key values in format [(john, 6)], so with the sum
	results := make([][]KeyValue, num_goroutines) // writing in specific index does not cause conflict
	for i:= range num_goroutines {
		wg3.Go( func() {
			i := i
			m := <-wordMaps
			toSend := make([]KeyValue, 0, len(m) ) // start point and capacity, append starts by filling form 0 up
			for key, value := range m{
				toSend = append(toSend, mr.wordCountReducer(key, value))
			}
			results[i] = toSend
		})
	}
	// 2. shuffle the output of count mapper to make (john, [1,1,1]) formats
	for range num_goroutines{
		wg2.Go( func() {
			wordMaps <- shuffle(<-lineWords) 
		})
	}
	// 1. get all single words out of string, does the word formatting (check if letter,...)
	for i := range num_goroutines {
		wg1.Go( func() {
			lineWords <- mr.wordCountMapper(input[i])
		})
	}
	// inverse creation order for waiting
	// go is not lazy and starts when the goroutine is defined, this simply forces it
	wg1.Wait()
	close(lineWords) // closes these channels lets wg2 finish, since it can stop listening
	wg2.Wait()
	close(wordMaps) // same here
	wg3.Wait()


	// convert my lists of unique values into a truly unique map
	result := map[string]int{}
	for _, pairs := range results {
		for _, pair := range pairs {
			result[pair.Key] += pair.Value
		}
	}
	return result
}
