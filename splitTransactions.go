// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Transaction struct {
	Amount  float64
	spender string
	lender  []string
}

func processTransactions(transactions []Transaction) map[string]map[string]float64 {
	debts := make(map[string]map[string]float64)
	for i, t := range transactions {
		splitAmount := t.Amount / float64(len(t.lender))
		fmt.Println("split amount for each transaction and i", splitAmount, i)

		for _, person := range t.lender {
			if t.spender == person {
				continue
			}

			if debts[t.spender] == nil {
				debts[t.spender] = make(map[string]float64)
			}
			debts[t.spender][person] += splitAmount
		}
	}
	return debts
}

func main() {
	transactions := []Transaction{
		{100, "A", []string{"A", "B", "C"}},
		{85.5, "C", []string{"A", "B"}},
	}

	debts := processTransactions(transactions)
	fmt.Println("Who owes Whom", debts)
	fmt.Println("Hello, 世界")
}
