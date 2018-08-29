package externals

import (
	"github.com/keybase/client/go/kbname"
	"github.com/keybase/client/go/libkb"
	keybase1 "github.com/keybase/client/go/protocol/keybase1"
)

func MakeAssertionContext() kbname.AssertionContext {
	return libkb.MakeAssertionContext(GetServices())
}

func NormalizeSocialAssertion(s string) (keybase1.SocialAssertion, bool) {
	return kbname.NormalizeSocialAssertion(MakeAssertionContext(), s)
}

func IsSocialAssertion(s string) bool {
	return kbname.IsSocialAssertion(MakeAssertionContext(), s)
}

func AssertionParseAndOnly(s string) (kbname.AssertionExpression, error) {
	return kbname.AssertionParseAndOnly(MakeAssertionContext(), s)
}

func AssertionParse(s string) (kbname.AssertionExpression, error) {
	return kbname.AssertionParse(MakeAssertionContext(), s)
}

func ParseAssertionsWithReaders(s string) (writers, readers []kbname.AssertionExpression, err error) {
	return kbname.ParseAssertionsWithReaders(MakeAssertionContext(), s)
}

func ParseAssertionList(s string) ([]kbname.AssertionExpression, error) {
	return kbname.ParseAssertionList(MakeAssertionContext(), s)
}
