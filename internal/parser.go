package internal

import (
	"amnesia-db/constants"
	"regexp"
	"strconv"
	"strings"
)

// OpCode ENUM
type OpCode int

const (
	GET OpCode = iota
	SET
	UPDATE
)

type Key string
type Value string

type Options map[OptionType]OptionValue

// OptionType ENUM
type OptionType int
type OptionValue struct {
	value    string
	intValue int32
}

const (
	TTL = iota
	NFETCH
)

type Query struct {
	Op      OpCode
	Key     Key
	Value   Value
	Options map[OptionType]OptionValue
}

func lvl1Tokenizer(s string) []string {
	var currTok = ""
	var tokens []string

	var ignoringLvl1 = false
	var ignoringLvl2 = false
	for _, char := range s {
		if char == ' ' && len(currTok) != 0 && !ignoringLvl1 && !ignoringLvl2 {
			if currTok == constants.WHERE {
				currTok = ""
				ignoringLvl2 = true
				continue
			}

			tokens = append(tokens, currTok)
			currTok = ""
		} else if char == '"' {
			ignoringLvl1 = !ignoringLvl1
		} else {
			currTok = currTok + string(char)
		}
	}

	if len(currTok) != 0 {
		tokens = append(tokens, currTok)
	}

	return tokens
}

func whereClauseParser(whereClause string) Options {
	var optionsRaw = strings.Split(whereClause, " ")
	var options Options = make(map[OptionType]OptionValue)

	for i := range optionsRaw {
		var optionType, optionValue = optionParser(optionsRaw[i])
		options[optionType] = optionValue
	}
	return options
}

var numberRegex = regexp.MustCompile("^[0-9]+")

func optionParser(rawOption string) (OptionType, OptionValue) {
	var lex = strings.Split(rawOption, "=")
	var oType = stringToOptionType(lex[0])
	var oValue OptionValue

	if oType == TTL {
		var n = numberRegex.FindString(lex[1])
		atoi, err := strconv.Atoi(n)
		if err != nil {
			panic("Something weird happened while doing atoi >> Panicking")
		}
		oValue = OptionValue{
			value:    lex[1][len(n):],
			intValue: int32(atoi),
		}
	}

	if oType == NFETCH {
		atoi, err := strconv.Atoi(lex[1])
		if err != nil {
			panic("Something weird happened while doing atoi >> Panicking")
		}
		oValue = OptionValue{
			intValue: int32(atoi),
		}
	}

	return oType, oValue
}

func stringToOptionType(s string) OptionType {
	if s == constants.TTL {
		return TTL
	} else if s == constants.NFETCH {
		return NFETCH
	}
	panic("Invalid Option Type Found! >> Panicking")
}

func ParseFromString(qs string) Query {
	var q Query
	var tokens = lvl1Tokenizer(qs)

	var op OpCode

	if tokens[0] == constants.SET {
		op = SET
	} else if tokens[0] == constants.GET {
		op = GET
	} else if tokens[0] == constants.UPDATE {
		op = UPDATE
	}

	q.Op = op

	key := Key(tokens[1])
	q.Key = key

	var nTokens = len(tokens)

	if nTokens >= 3 {
		if tokens[2] != constants.AS {
			panic("Invalid Query! >> Panicking")
		}
	}

	if nTokens >= 4 {
		q.Value = Value(tokens[3])
	}

	if nTokens >= 5 {
		q.Options = whereClauseParser(tokens[4])
	}

	return q
}
