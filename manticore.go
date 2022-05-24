package manticore

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

const (
	SymbolPlaceholder = '?'
	SymbolBackslash   = '\\'
	SymbolSingleQuote = '\''
	SymbolDoubleQuote = '"'
)

func MapRows(rows *sql.Rows, cols ...Column) (map[string]interface{}, error) {
	return nil, nil
}

// Escape is a blaintant copy from https://stackoverflow.com/a/35257216/772008
// along with all the charactuers that need to be esacped according to the
// Manticoresearch Documentation at
// https://manual.manticoresearch.com/Searching/Full_text_matching/Escaping#Escaping-characters-in-query-string
// This is used to try and combat the fact that Manticore + MySQL drivers do
// not operate properly using the binary format. In other words: no
// parameterization on queries, nor prepared statements. If generating SQL by
// hand, this should be used on the _ARGS_, and not the SQL query as a whole.
func Escape(sql string) string {
	dest := make([]byte, 0, 2*len(sql))
	for i := 0; i < len(sql); i++ {
		c := sql[i]
		escape := 0

		switch c {
		case 0: // Must be escaped for 'mysql'
			c = '0'
			escape = 1
		case '\n': // Must be escaped for logs
			c = 'n'
			escape = 1
		case '\r':
			c = 'r'
			escape = 1
		case '\032': //十进制26,八进制32,十六进制1a, /* This gives problems on Win32 */
			c = 'Z'
			escape = 1
		case '\'':
			escape = 1
		case '!', '"', '$', '(', ')', '-', '/', '<', '@', '^', '~', '&':
			escape = 2
		case '\\':
			escape = 3
		}

		switch escape {
		case 1:
			dest = append(dest, '\\', c)
		case 2:
			dest = append(dest, '\\', '\\', c)
		case 3:
			dest = append(dest, '\\', '\\', '\\', c)
		default:
			dest = append(dest, c)
		}
	}

	return string(dest)
}

func queryBuilder(query string, args ...any) (string, error) {
	// fmt.Printf("query: %#v\targs: %#v\n", query, args)
	maxArgs := len(args)
	argIndex := 0
	sb := strings.Builder{}
	sb.Grow(len(query) * 2)
	var prev rune
	for _, r := range []rune(query) {
		// only care about ?
		if r != SymbolPlaceholder {
			sb.WriteRune(r)
			prev = r
			continue
		}

		// make sure we're not at \? in the query string; escaped question marks
		// are allowed in the query string.
		if prev == SymbolBackslash {
			sb.WriteRune(r)
			prev = r
			continue
		}

		if argIndex >= maxArgs {
			return "", fmt.Errorf("expected %d args; have at least %d", maxArgs, argIndex+1)
		}
		arg := args[argIndex]

		val, err := paramToString(arg)
		if err != nil {
			return "", err
		}

		sb.WriteString(val)
		argIndex++
		prev = r
	}

	return sb.String(), nil
}

func paramToString(v any) (val string, err error) {
	switch v.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		val = fmt.Sprintf("%v", v)
	case float32, float64:
		val = fmt.Sprintf("%0.04f", v)
	case time.Time:
		ts := v.(time.Time)
		val = fmt.Sprintf("%d", ts.Unix())
	case bool:
		n := 0
		if b, ok := v.(bool); ok && b {
			n = 1
		}
		val = fmt.Sprintf("%v", n)
	default:
		// default to string! if casting fails, throw error
		if _, ok := v.(string); !ok {
			err = fmt.Errorf("cannot parse type: %T %#v", v, v)
			return
		}
		val = v.(string)
	}

	return
}
