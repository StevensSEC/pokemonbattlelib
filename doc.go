// pokemonbattlelib is a library for simulating pokemon battles using rules up to generation 4 (currently).
//
// Battles are represented as the result of a sequence of Transactions. Transactions define what state needs to change,
// and can queue up subsequent transactions via Mutate. Agents control Parties. Each round, each active (on the battlefield)
// Pokemon gets one Turn. The Agent that controls the Party of the Pokemon dictates what Turn the Pokemon will make.
// Then, the Turns queue up Transactions accordingly.
package pokemonbattlelib
