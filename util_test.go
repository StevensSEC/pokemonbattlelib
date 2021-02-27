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
	return a.NatDex == b.NatDex &&
		a.Nature.name == b.Nature.name &&
		a.Gender == b.Gender &&
		a.Elemental == b.Elemental
}

// Used for custom gomega matchers. For simplicity, the fields of the struct are hardcoded. If we need to add more fields to `target` something is probably wrong.
func compareTargets(a, b target) bool {
	return comparePokemon(&a.Pokemon, &b.Pokemon) &&
		a.party == b.party &&
		a.partySlot == b.partySlot &&
		a.Team == b.Team
}

// Used for custom gomega matchers.
func compareTransaction(a, b Transaction) bool {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return false
	}

	rA := reflect.ValueOf(a)
	rB := reflect.ValueOf(b)

	fieldsMatch := true
	for i := 0; i < rA.NumField(); i++ {
		typeField := rA.Type().Field(i)
		rfA := rA.Field(i)
		rfB := rB.FieldByName(typeField.Name)

		if rfA.Type() == reflect.TypeOf(&Pokemon{}) {
			a := rfA.Interface().(*Pokemon)
			b := rfB.Interface().(*Pokemon)
			if !comparePokemon(a, b) {
				fieldsMatch = false
				break
			}
			continue
		} else if rfA.Type() == reflect.TypeOf(target{}) {
			a := rfA.Interface().(target)
			b := rfB.Interface().(target)
			if !compareTargets(a, b) {
				fieldsMatch = false
				break
			}
			continue
		}

		if !reflect.DeepEqual(rfA.Interface(), rfB.Interface()) {
			fieldsMatch = false
			break
		}
	}

	return fieldsMatch
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
	return fmt.Sprintf("Expected the sequence of transactions to include: %T",
		matcher.expected,
	)
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
		seq = append(seq, fmt.Sprintf("%d: %T", i, t))
	}
	return fmt.Sprintf("Expected the sequence of transactions to have these transactions in this order:\n%s",
		strings.Join(seq, "\n"),
	)
}

func (matcher *orderedTransactionMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	seq := []string{}
	for i, t := range matcher.expected {
		seq = append(seq, fmt.Sprintf("%d: %T", i, t))
	}
	return fmt.Sprintf("Expected the sequence of transactions to NOT have these transactions in this order:\n%s",
		strings.Join(seq, "\n"),
	)
}
