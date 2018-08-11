package ql

func (l List) Stmts() []stmt {
	return l.l
}

type FromSources interface {
	Sources() []string
}

func (s *selectStmt) From() FromSources {
	return s.from
}

func (j *joinRset) Sources() []string {
	a := make([]string, 0, len(j.sources))
	for _, pair0 := range j.sources {
		pair := pair0.([]interface{})
		switch x := pair[0].(type) {
		case string: // table name
			a = append(a, x)
		case *selectStmt:
			a = append(a, x.From().Sources()...)
		default:
			panic("internal error 054")
		}
	}
	return a
}

type emptyMutex struct{}

func (m *emptyMutex) Lock()   {}
func (m *emptyMutex) Unlock() {}
