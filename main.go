package main

import "fmt"

type tokenType uint8

const (
	group           tokenType = iota
	bracket         tokenType = iota
	or              tokenType = iota
	repeat          tokenType = iota
	literal         tokenType = iota
	groupUncaptured tokenType = iota
)

type token struct {
	tokenType tokenType
	value     interface{}
}

type parseContext struct {
	position int
	tokens   []token
}

func parse(regex string) *parseContext {
	ctx := &parseContext{
		position: 0,
		tokens:   []token{},
	}
	for ctx.position < len(regex) {
		process(regex, ctx)
		ctx.position++
	}
	return ctx
}

func parseGroup(regex string, ctx *parseContext) {
	ctx.position += 1
	for regex[ctx.position] != ')' {
		process(regex, ctx)
		ctx.position += 1
	}
}

func parseBracket(regex string, ctx *parseContext) {
	ctx.position++
	var literals []string
	for regex[ctx.position] != ']' {
		ch := regex[ctx.position]
		if ch == '-' {
			next := regex[ctx.position+1]
			prev := literals[len(literals)-1][0]
			literals[len(literals)-1] = fmt.Sprintf("%c%c", prev, next)
		} else {
			literals = append(literals, fmt.Sprintf("%c", ch))
		}
		ctx.position++
	}
	literalsSet := map[uint8]bool{}
	for _, l := range literals { // <6>
		for i := l[0]; i <= l[len(l)-1]; i++ { // <7>
			literalsSet[i] = true
		}
	}

	ctx.tokens = append(ctx.tokens, token{ // <8>
		tokenType: bracket,
		value:     literalsSet,
	})
}

func parseOr()              {}
func parseRepeat()          {}
func parseRepeatSpecified() {}
func process(regex string, ctx *parseContext) {
	ch := regex[ctx.position]
	if ch == '(' { // <1>
		groupCtx := &parseContext{
			position: ctx.position,
			tokens:   []token{},
		}
		parseGroup(regex, groupCtx)
		ctx.tokens = append(ctx.tokens, token{
			tokenType: group,
			value:     groupCtx.tokens,
		})
	} else if ch == '[' { // <2>
		parseBracket(regex, ctx)
	} else if ch == '|' { // <3>
		parseOr(regex, ctx)
	} else if ch == '*' || ch == '?' || ch == '+' { // <4>
		parseRepeat(regex, ctx)
	} else if ch == '{' { // <5>
		parseRepeatSpecified(regex, ctx)
	} else { // <6>
		// literal
		t := token{
			tokenType: literal,
			value:     ch,
		}

		ctx.tokens = append(ctx.tokens, t)
	}
}
