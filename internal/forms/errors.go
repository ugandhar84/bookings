package forms

type errors map[string][]string

// Add adds the error message
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)

}

// Get return the first error message
func (e errors) Get(field string) string {
	es := e[field]

	if len(es) == 0 {
		return field
	}
	return es[0]
}
