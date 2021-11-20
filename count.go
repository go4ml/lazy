package lazy

func (zf Source) Count() (count int, err error) {
	err = zf.Drain(func(int)[]Worker{
		return []Worker{func(_ int, v interface{}, _ error) (_ error) {
			if v != nil {
				count++
			}
			return
		}}
	})
	return
}

func (zf Source) MustCount() int {
	count, err := zf.Count()
	if err != nil {
		panic(err)
	}
	return count
}
