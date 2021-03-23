package pokemonbattlelib

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/onsi/gomega/types"
)

// Gomega custom matchers, see: https://onsi.github.io/gomega/#adding-your-own-matchers

// Used for custom gomega matchers. Checks to see if a is probably the same pokemon as b based on values that are unlikely to change as a result of a transaction.
func comparePokemon(a, b *Pokemon) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	}
	return a.NatDex == b.NatDex &&
		a.Nature == b.Nature &&
		a.Gender == b.Gender &&
		a.Type == b.Type
}

// Used for custom gomega matchers. For simplicity, the fields of the struct are hardcoded. If we need to add more fields to `target` something is probably wrong.
func compareTargets(a, b target) bool {
	return comparePokemon(&a.Pokemon, &b.Pokemon) &&
		a.party == b.party &&
		a.partySlot == b.partySlot &&
		a.Team == b.Team
}

// Helper struct for finding differences in objects for testing
type diff struct {
	expected interface{} // The expected value
	got      interface{} // The actual value received
}

// Creates a diff of expected fields vs received fields in same-type transactions
// Also returns the number of fields that were compared
func transactionDiff(expected, got Transaction) map[string]diff {
	// types that must match exactly, including zeroed values
	exactMatchTypes := []reflect.Type{
		reflect.TypeOf(true),
		reflect.TypeOf(FailMiss),
		reflect.TypeOf(WeatherClearSkies),
		reflect.TypeOf(StatusNone),
	}

	result := make(map[string]diff)
	rA := reflect.ValueOf(expected)
	rB := reflect.ValueOf(got)
	for i := 0; i < rA.NumField(); i++ {
		typeField := rA.Type().Field(i)
		rfA := rA.Field(i)
		rfB := rB.FieldByName(typeField.Name)

		if rfA.Type() == reflect.TypeOf(&Pokemon{}) {
			a := rfA.Interface().(*Pokemon)
			b := rfB.Interface().(*Pokemon)
			if !comparePokemon(a, b) {
				result[typeField.Name] = diff{
					expected: a,
					got:      b,
				}
			}
			continue
		} else if rfA.Type() == reflect.TypeOf(target{}) {
			a := rfA.Interface().(target)
			b := rfB.Interface().(target)
			if !compareTargets(a, b) {
				result[typeField.Name] = diff{
					expected: a,
					got:      b,
				}
			}
			continue
		} else if rfA.Kind() == reflect.Struct || rfA.Kind() == reflect.Interface || rfA.Kind() == reflect.Map || rfA.Kind() == reflect.Array || rfA.Kind() == reflect.Slice {
			if !reflect.DeepEqual(rfA.Interface(), rfB.Interface()) {
				result[typeField.Name] = diff{
					expected: rfA.Interface(),
					got:      rfB.Interface(),
				}
			}
		} else {
			// Special case to allow fields with primitive types or nil pointers to be ignored when comparing.
			// If either A or B is a type's zero value, or nil, it won't bother comparing them.
			// For example, you should be able to omit Damage from the expected DamageTransaction if you don't want to check that.
			// Example use case: Expect(transactions).ToNot(HaveTransaction(DamageTransaction{}))

			mustExactMatch := false
			for _, t := range exactMatchTypes {
				if t == rfA.Type() {
					mustExactMatch = true
					break
				}
			}

			if rfA.Kind() == reflect.Ptr {
				if !rfA.IsNil() && !rfB.IsNil() {
					if !reflect.DeepEqual(rfA.Interface(), rfB.Interface()) {
						result[typeField.Name] = diff{
							expected: rfA.Interface(),
							got:      rfB.Interface(),
						}
					}
				}
			} else if mustExactMatch {
				if !reflect.DeepEqual(rfA.Interface(), rfB.Interface()) {
					result[typeField.Name] = diff{
						expected: rfA.Interface(),
						got:      rfB.Interface(),
					}
				}
			} else if rfA.Interface() != reflect.Zero(rfA.Type()).Interface() && rfB.Interface() != reflect.Zero(rfB.Type()).Interface() {
				if !reflect.DeepEqual(rfA.Interface(), rfB.Interface()) {
					result[typeField.Name] = diff{
						expected: rfA.Interface(),
						got:      rfB.Interface(),
					}
				}
			}
		}
	}
	return result
}

// Used for custom gomega matchers.
func compareTransaction(a, b Transaction) bool {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return false
	}
	return len(transactionDiff(a, b)) == 0
}

// Used for custom gomega matchers for failure messages. Uses reflection to find the index of the first
// transaction with a matching type, and how many transactions match the type. Returns (-1, 0) if not found.
func findCountTransactionIdxWithMatchingType(transactions []Transaction, a Transaction) (first int, count int) {
	first = -1
	for i, t := range transactions {
		if reflect.TypeOf(t) == reflect.TypeOf(a) {
			count++
			if first == -1 {
				first = i
			}
		}
	}
	return first, count
}

// Gets the closest transaction to the desired one, and return a diff in fields
func getClosestTransaction(check []Transaction, want Transaction) map[string]diff {
	var best map[string]diff
	bestDiff := 999
	for _, t := range check {
		if reflect.TypeOf(t) != reflect.TypeOf(want) {
			continue
		}
		result := transactionDiff(want, t)
		if len(result) < bestDiff {
			bestDiff = len(result)
			best = result
		}
	}
	return best
}

// Given a sequence of transactions, match if a given transaction is present in the sequence.
type singleTransactionMatcher struct {
	expected Transaction
}

// Check to see if this transaction occured.
func HaveTransaction(expected Transaction) types.GomegaMatcher {
	return &singleTransactionMatcher{
		expected: expected,
	}
}

func (matcher *singleTransactionMatcher) Match(actual interface{}) (success bool, err error) {
	switch transactions := actual.(type) {
	case []Transaction:
		for _, t := range transactions {
			if compareTransaction(t, matcher.expected) {
				return true, nil
			}
		}
		return false, nil
	default:
		return false, errors.New("Was not given a []Transaction")
	}
}

func (matcher *singleTransactionMatcher) FailureMessage(actual interface{}) (message string) {
	switch transactions := actual.(type) {
	case []Transaction:
		first, count := findCountTransactionIdxWithMatchingType(transactions, matcher.expected)
		if first == -1 {
			got := ""
			for _, t := range transactions {
				got += fmt.Sprintf("- %T\n", t)
			}
			return fmt.Sprintf("Expected the sequence of transactions to include: %T, but received the following transactions:\n%s",
				matcher.expected,
				got,
			)
		} else if count == 1 {
			diffText := ""
			result := transactionDiff(matcher.expected, transactions[first])
			total := reflect.ValueOf(matcher.expected).NumField()
			for name, val := range result {
				diffText += fmt.Sprintf("- %s\nExpected: %v\nReceived: %v\n", name, val.expected, val.got)
			}
			return fmt.Sprintf("%d/%d fields of the %T do not match:\n%s",
				len(result), total, matcher.expected, diffText)
		} else {
			diffText := ""
			closest := getClosestTransaction(transactions, matcher.expected)
			total := reflect.ValueOf(matcher.expected).NumField()
			for name, val := range closest {
				diffText += fmt.Sprintf("- %s\nExpected: %v\nReceived: %v\n", name, val.expected, val.got)
			}
			return fmt.Sprintf("Closest %T (%d total) has %d/%d fields that don't match:\n%s",
				matcher.expected, count, len(closest), total, diffText)
		}
	default:
		return fmt.Sprintf("Actual's type is %T, when it should be []Transaction", actual)
	}
}

func (matcher *singleTransactionMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected the sequence of transactions NOT to include: %T",
		matcher.expected,
	)
}

// Given a sequence of transactions, match if a given set of transactions is present in the sequence, and the order matches.
type orderedTransactionMatcher struct {
	expected []Transaction
}

// Check to see if these transactions occured in this order.
func HaveTransactionsInOrder(expected ...Transaction) types.GomegaMatcher {
	return &orderedTransactionMatcher{
		expected: expected,
	}
}

func (matcher *orderedTransactionMatcher) Match(actual interface{}) (success bool, err error) {
	switch transactions := actual.(type) {
	case []Transaction:
		expectedIdx := 0
		for _, t := range transactions {
			if compareTransaction(t, matcher.expected[expectedIdx]) {
				expectedIdx++
				if expectedIdx == len(matcher.expected) {
					break
				}
			}
		}
		return expectedIdx == len(matcher.expected), nil
	default:
		return false, errors.New("Was not given a []Transaction")
	}
}

func (matcher *orderedTransactionMatcher) FailureMessage(actual interface{}) (message string) {
	seq := []string{}
	for i, t := range matcher.expected {
		seq = append(seq, fmt.Sprintf("%d: %T: %+v", i, t, t))
	}
	return fmt.Sprintf("Expected the sequence of transactions to have these transactions in this order:\n%s",
		strings.Join(seq, "\n"),
	)
}

func (matcher *orderedTransactionMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	seq := []string{}
	for i, t := range matcher.expected {
		seq = append(seq, fmt.Sprintf("%d: %T: %+v", i, t, t))
	}
	return fmt.Sprintf("Expected the sequence of transactions to NOT have these transactions in this order:\n%s",
		strings.Join(seq, "\n"),
	)
}

// Tools for testing the library
// Custom RNG struct which allows for predictable RNG output in a battle
type TestRNG struct {
	rolls  []bool
	rounds int
}

func (g *TestRNG) SetSeed(uint) {}
func (g *TestRNG) Get(min, max int) int {
	return max
}
func (g *TestRNG) Roll(x, y int) bool {
	v := g.rolls[g.rounds%len(g.rolls)]
	g.rounds += 1
	return v
}

// Never rolls random effects
func NeverRNG() *TestRNG {
	return &TestRNG{rolls: []bool{false}}
}

// Always rolls random effects
func AlwaysRNG() *TestRNG {
	return &TestRNG{rolls: []bool{true}}
}

// Always hit, never crit
func SimpleRNG() *TestRNG {
	return &TestRNG{rolls: []bool{true, false}}
}
