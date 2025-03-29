package main

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

func process(regex string, ctx *parseContext) {

}
