// Generated by: main.exe
// TypeWriter: stringer
// Directive: +gen on TokenType

package internals

import (
	"fmt"
)

const _TokenType_name = "TokenTypeNoneTokenTypeEndOfInputTokenTypeRangeInclusiveTokenTypeRangeHalfOpenTokenTypeIntervalTokenTypeNotTokenTypeOpenParenTokenTypeCloseParenTokenTypeOpenCurlyTokenTypeCloseCurlyTokenTypeForwardSlashTokenTypeCommaTokenTypeWildcardTokenTypePositiveIntegerTokenTypeNegativeIntegerTokenTypeExpressionNameTokenTypeDayLiteral"

var _TokenType_index = [...]uint16{0, 13, 32, 55, 77, 94, 106, 124, 143, 161, 180, 201, 215, 232, 256, 280, 303, 322}

func (i TokenType) String() string {
	if i < 0 || i+1 >= TokenType(len(_TokenType_index)) {
		return fmt.Sprintf("TokenType(%d)", i)
	}
	return _TokenType_name[_TokenType_index[i]:_TokenType_index[i+1]]
}
