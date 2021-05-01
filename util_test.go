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
	return a == b
}

// Used for custom gomega matchers. For simplicity, the fields of the struct are hardcoded. If we need to add more fields to `target` something is probably wrong.
func compareTargets(a, b target) bool {
	return comparePokemon(a.Pokemon, b.Pokemon) &&
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

		if !rfB.IsValid() {
			result[typeField.Name] = diff{
				expected: rfA.Interface(),
				got:      "invalid reflection value",
			}
			continue
		}

		if rfA.Type() != rfB.Type() {
			result[typeField.Name] = diff{
				expected: rfA.Type(),
				got:      rfB.Type(),
			}
			continue
		}

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
		} else if rfA.Type() == reflect.TypeOf(&Move{}) {
			if !rfA.IsNil() && !rfB.IsNil() {
				mvA := rfA.Interface().(*Move)
				mvB := rfB.Interface().(*Move)
				if mvA.Id != mvB.Id {
					result[typeField.Name] = diff{
						expected: mvA.String(),
						got:      mvB.String(),
					}
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

// Creates a text representation of the diff result of transactions
func getDiffText(check []Transaction, want Transaction) string {
	diffText := ""
	closest := getClosestTransaction(check, want)
	total := reflect.ValueOf(want).NumField()
	for name, val := range closest {
		diffText += fmt.Sprintf("- %s\nExpected: %v\nReceived: %v\n", name, val.expected, val.got)
	}
	return fmt.Sprintf("%T has %d/%d fields that do not match:\n%s", want, len(closest), total, diffText)
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
			gotText := ""
			for _, t := range transactions {
				gotText += fmt.Sprintf("- %T\n", t)
			}
			return fmt.Sprintf("Expected the sequence of transactions to include %T, but received:\n%s",
				matcher.expected,
				gotText,
			)
		} else if count == 1 {
			return getDiffText(transactions, matcher.expected)
		} else {
			diffText := getDiffText(transactions, matcher.expected)
			return fmt.Sprintf("The closest of %d total %s", count, diffText)
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

func orderedTransactionDiffLine(idx int, t Transaction) string {
	line := fmt.Sprintf("%d. %T", idx+1, t)
	switch tt := t.(type) {
	case UseMoveTransaction:
		line += fmt.Sprintf(" - User [%d,%d] Target: [%d,%d] Move: %s",
			tt.User.party, tt.User.partySlot,
			tt.Target.party, tt.Target.partySlot,
			tt.Move,
		)
	case DamageTransaction:
		line += fmt.Sprintf(" - Target: [%d, %d] Damage: %d",
			tt.Target.party, tt.Target.partySlot,
			tt.Damage,
		)
	case InflictStatusTransaction:
		line += fmt.Sprintf(" - Target: %s Status: %s",
			tt.Target,
			tt.StatusEffect)
	}
	line += "\n"
	return line
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

func checkTransactionOrder(check, want []Transaction) (success bool, diffText string) {
	i := 0
	fails := make([]Transaction, 0)
	for _, t := range check {
		if i == len(want) {
			break
		}
		if reflect.TypeOf(t) != reflect.TypeOf(want[i]) {
			continue
		}
		d := transactionDiff(want[i], t)
		if len(d) == 0 {
			i += 1
			fails = make([]Transaction, 0)
		} else {
			fails = append(fails, t)
		}
	}
	if i == len(want) {
		return true, ""
	}
	diffText = getDiffText(fails, want[i])
	return false, diffText
}

func (matcher *orderedTransactionMatcher) Match(actual interface{}) (success bool, err error) {
	switch t := actual.(type) {
	case []Transaction:
		if len(t) < len(matcher.expected) {
			return false, fmt.Errorf("Received %d, but expected at least %d transactions.", len(matcher.expected), len(t))
		}
		result, _ := checkTransactionOrder(t, matcher.expected)
		return result, nil
	default:
		return false, errors.New("Was not given a []Transaction")
	}
}

func (matcher *orderedTransactionMatcher) FailureMessage(actual interface{}) (message string) {
	wantOrder := ""
	for i, t := range matcher.expected {
		wantOrder += orderedTransactionDiffLine(i, t)
	}
	switch transactions := actual.(type) {
	case []Transaction:
		gotOrder := ""
		for i, t := range transactions {
			gotOrder += orderedTransactionDiffLine(i, t)
		}
		msg := fmt.Sprintf("Expected the sequence of transactions to have these transactions in this order:\n%s"+
			"\nReceived the following transactions:\n%s",
			wantOrder, gotOrder,
		)
		_, result := checkTransactionOrder(transactions, matcher.expected)
		return fmt.Sprintf("%s\nThe closest transaction that failed to match is shown below:\n%s", msg, result)
	default:
		return fmt.Sprintf("Actual's type is %T, when it should be []Transaction", actual)
	}
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

/* Tools for testing the library */
// Check for damage dealt (if any) by a Pokemon in battle
func DamageDealt(t []Transaction, p *Pokemon) int {
	var usemove *UseMoveTransaction
	var start int
	for i := start; i < len(t); i++ {
		if v, ok := t[i].(UseMoveTransaction); !ok {
			continue
		} else if v.User.Pokemon == p {
			usemove = &v
			break
		}
	}
	if usemove == nil {
		// This might work better as a gomega matcher
		return -1 // failure, pokemon did not use a move
	}

	for i := start; i < len(t); i++ {
		if d, ok := t[i].(DamageTransaction); ok {
			if d.Move == usemove.Move {
				return int(d.Damage)
			}
		}
	}
	return -2 // failure, move was used but did not deal damage
}

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

// Create a single Battle. A single battle is a battle where only 1 pokemon is sent out at a time, and there can only be 2 parties.
// This function should only be used in tests.
func NewSingleBattle(p1 *Party, a1 *Agent, p2 *Party, a2 *Agent) *Battle {
	b := NewBattle()
	b.AddParty(p1, a1, 0)
	b.AddParty(p2, a2, 1)
	return b
}

// Create a single Battle with only 1 pokemon in each party.
// This function should only be used in tests.
func New1v1Battle(p1 *Pokemon, a1 *Agent, p2 *Pokemon, a2 *Agent) *Battle {
	return NewSingleBattle(NewOccupiedParty(p1), a1, NewOccupiedParty(p2), a2)
}

// Deprecated: Creates a new party to store Pokemon and assigns them to a team
// This function should only be used in tests.
func newBattlePartyOld(agent *Agent, team int) *battleParty {
	return &battleParty{
		Party:         NewParty(),
		Agent:         agent,
		activePokemon: make(map[int]*Pokemon),
		team:          team,
	}
}

// Creates a new party and fills it out with the passed Pokemon
// This function should only be used in tests.
func newOccupiedBattleParty(agent *Agent, team int, pkmn ...*Pokemon) *battleParty {
	party := newBattlePartyOld(agent, team)
	err := party.AddPokemon(pkmn...)
	if err != nil {
		panic(err)
	}
	return party
}
