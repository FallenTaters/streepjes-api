package wherebuilder

type WhereBuilder struct {
	begin  string
	wheres []string
}

func New(begin string) WhereBuilder {
	return WhereBuilder{begin: begin}
}

func (wb *WhereBuilder) Where(where string) {
	wb.wheres = append(wb.wheres, where)
}

func (wb WhereBuilder) Query() string {
	q := wb.begin
	if len(wb.wheres) == 0 {
		return q + `;`
	}

	q += ` WHERE `
	and := false
	for _, where := range wb.wheres {
		if and {
			q += ` AND `
		}
		and = true

		q += where
	}

	return q + `;`
}
