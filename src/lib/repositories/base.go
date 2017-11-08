package repositories

// Base application repository.
type base struct{}

func (r *base) Insert(v interface{}) error {
	return nil
}

func (r *base) Update(v interface{}) error {
	return nil
}

func (r *base) Remove(v interface{}) error {
	return nil
}
