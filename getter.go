package getters

type Getter interface {
	Get(key interface{}) (value interface{}, ok bool)
}

type MultipleGetter []Getter

func (this MultipleGetter) Get(key interface{}) (value interface{}, ok bool) {
	for _, g := range this {
		if value, ok = g.Get(key); ok {
			return
		}
	}
	return
}

func (this *MultipleGetter) Append(g ...Getter) *MultipleGetter {
	(*this) = append((*this), g...)
	return this
}

func (this *MultipleGetter) Prepend(g ...Getter) *MultipleGetter {
	(*this) = append(g[:], (*this)...)
	return this
}

type Func func(key interface{}) (value interface{}, ok bool)

func (f Func) Get(key interface{}) (value interface{}, ok bool) {
	return f(key)
}

func New(getter func(key interface{}) (value interface{}, ok bool)) Getter {
	return Func(getter)
}
