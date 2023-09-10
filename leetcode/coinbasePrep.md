
**Technical interview**
Design an iterator which can output infinity even number
two api: next/hasNext
1,3,5,7,...,inf
Iterators:

Custom Iterator (https://leetcode.com/problems/peeking-iterator/) 
Odd, even having infinite stream (Even takes odd as dependency)
Negative Iterator( takes a dependency iterator)
InterLeavingIterator (takes a list of iterators)
Limit Iterator (takes a limit and terminates the infinite iterator)

currency exchange https://www.1point3acres.com/bbs/thread-788575-1-1.html

connect 4 game 

In Memory File System 

**system design**
- Credit fraud check system
- design slack and slack related features (show "some one is typing" etc)
- time series database
 

- Distributed Systems: Understanding the concepts of distributed systems, including scalability, fault tolerance, consistency, and availability.

- System Architecture: Designing the high-level architecture of a system, including components, their interactions, and communication protocols.

- Database Design: Designing database schemas, choosing appropriate database technologies, and understanding data modeling concepts.

- Caching: Knowledge of caching mechanisms to improve system performance, such as in-memory caches or content delivery networks (CDNs).

- Load Balancing: Understanding load balancing techniques to distribute traffic across multiple servers or instances.

- Message Queues: Knowledge of message queuing systems for asynchronous communication between components.


// CURRENCY CONVERSION 
// https://interviewing.io/questions/currency-conversion
// Function to calculate conversion rates
func calculateConversionRates(rates [][]interface{}, queries [][]string) []float64 {
	// Build graph
	graph := make(map[string]map[string]float64)
	for _, rate := range rates {
		fromCurrency := rate[0].(string)
		toCurrency := rate[1].(string)
		value := rate[2].(float64)

		if _, exists := graph[fromCurrency]; !exists {
			graph[fromCurrency] = make(map[string]float64)
		}
		if _, exists := graph[toCurrency]; !exists {
			graph[toCurrency] = make(map[string]float64)
		}

		graph[fromCurrency][toCurrency] = value
		graph[toCurrency][fromCurrency] = 1.0 / value
	}

	// Perform DFS for each query
	var result []float64
	for _, query := range queries {
		fromCurrency := query[0]
		toCurrency := query[1]

		visited := make(map[string]bool)
		rate := dfs(graph, fromCurrency, toCurrency, 1.0, visited)
		result = append(result, rate)
	}

	return result
}

// Depth-first search to find conversion rate
func dfs(graph map[string]map[string]float64, start, end string, value float64, visited map[string]bool) float64 {
	if _, exists := graph[start]; !exists || visited[start] {
		return -1.0
	}

	if start == end {
		return value
	}

	visited[start] = true
	neighbors := graph[start]
	for neighbor, neighborValue := range neighbors {
		rate := dfs(graph, neighbor, end, value*neighborValue, visited)
		if rate != -1.0 {
			return rate
		}
	}

	return -1.0
}

func main() {
	rates := [][]interface{}{{"USD", "JPY", 100.0}, {"JPY", "CHN", 20.0}, {"CHN", "THAI", 200.0}}
	queries := [][]string{{"USD", "CHN"}, {"JPY", "THAI"}, {"USD", "AUD"}}

	result := calculateConversionRates(rates, queries)
	fmt.Println(result)
}

