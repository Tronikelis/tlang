package ast

import (
	"errors"
	"fmt"
	"tlang/src/tokens"
)

type Traverser struct {
	tokens  []tokens.Token
	pointer int
}

func NewTraverser(tokens []tokens.Token) Traverser {
	return Traverser{
		tokens: tokens,
	}
}

func (t *Traverser) tokenAt(index int) (tokens.Token, bool) {
	if index >= len(t.tokens)-1 || index < 0 {
		return tokens.Token{}, false
	}

	return t.tokens[index], true
}

func (t *Traverser) Next() (tokens.Token, bool) {
	token, exists := t.tokenAt(t.pointer)
	if !exists {
		return token, false
	}

	t.pointer += 1

	return token, true
}

func (t *Traverser) PeekNext() (tokens.Token, bool) {
	t.pointer += 1
	defer func() {
		t.pointer -= 1
	}()

	token, exists := t.Next()

	return token, exists
}

func (t *Traverser) Expect(ttype tokens.TokenType) error {
	if token, exists := t.tokenAt(t.pointer); exists {
		if token.Type != ttype {
			return errors.New(fmt.Sprintf("expected %v, got %v", ttype, token.Type))
		}
	}

	return nil
}

func (t *Traverser) ExpectNext(ttype tokens.TokenType) error {
	t.pointer += 1
	defer func() {
		t.pointer -= 1
	}()

	return t.Expect(ttype)
}

func (t *Traverser) ExpectNextOneOf(ttypes []tokens.TokenType) error {
	t.pointer += 1
	defer func() {
		t.pointer -= 1
	}()

	found := false

	for _, ttype := range ttypes {
		if err := t.Expect(ttype); err == nil {
			found = true
			break
		}
	}

	if found {
		return nil
	}

	return errors.New("TODO: expected one of many")
}
