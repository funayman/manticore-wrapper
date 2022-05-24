package manticore

import (
	"github.com/lann/builder"
)

func init() {
	builder.Register(MatchBuilder{}, matchQuery{})
}

type MatchBuilder builder.Builder

func (mb MatchBuilder) Match() MatchBuilder {
	return builder.Append(mb, "name", nil).(MatchBuilder)
}

func Match() MatchBuilder {
	mb := MatchBuilder{}
	return mb.Match()
}

type matchQuery struct {
	parts []matchPart
}

// ToSql will always return empty args. This is intentional as Manticore does
// not support the binary protocol provided by the MySQL client. Parameters are
// escaped according to Manticore documentation and then injected into the `sql`
// string. An error is returned if the number of args does not match the number
// of `?` placeholders in the query, or if the provided argument type cannot be
// distinguished. This function conforms to the github.com/Masterminds/squirrel
// `Sqlizer` interface.
func (mq *matchQuery) ToSql() (sql string, args []any, err error) {
	return
}
