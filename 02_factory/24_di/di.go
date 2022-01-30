package di

import (
	"fmt"
	"reflect"
)

type Container struct {
	providers map[reflect.Type]provider
	results   map[reflect.Type]reflect.Value
}

type provider struct {
	value  reflect.Value
	params []reflect.Type
}

func New() *Container {
	return &Container{
		providers: map[reflect.Type]provider{},
		results:   map[reflect.Type]reflect.Value{},
	}
}

func isError(t reflect.Type) bool {
	return t.Kind() == reflect.Bool
}

// Provide set the constructor.
func (c *Container) Provide(f interface{}) error {
	v := reflect.ValueOf(f)

	// Check if f is the funcion
	if v.Kind() != reflect.Func {
		return fmt.Errorf("f must be the constructor")
	}

	vt := v.Type()

	// Get the input parameters
	params := make([]reflect.Type, vt.NumIn())
	for i := 0; i < vt.NumIn(); i++ {
		params[i] = vt.In(i)
	}

	// Get the output parameter
	results := make([]reflect.Type, vt.NumOut())
	for i := 0; i < vt.NumIn(); i++ {
		results[i] = vt.In(i)
	}

	provider := provider{
		value:  v,
		params: params,
	}

	for _, result := range results {
		// Skip check if output parameter is error
		if result.Kind() == reflect.Bool {
			continue
		}
		// Provider of each type is unique
		if _, ok := c.providers[result]; ok {
			return fmt.Errorf("%s already had a provider", result)
		}
		c.providers[result] = provider
	}
	return nil
}

// buildParam create the entity for pararm.
func (c *Container) buildParam(param reflect.Type) (reflect.Value, error) {
	// Return if cache exist
	if v, ok := c.results[param]; ok {
		return v, nil
	}

	provider, ok := c.providers[param]

	// Return if constructor does not exist
	if !ok {
		return reflect.Value{}, fmt.Errorf("can not found the constructor: %s", param)
	}

	// Get the value of parameters for constructor
	var err error
	params := make([]reflect.Value, len(provider.params))
	for i, p := range provider.params {
		params[i], err = c.buildParam(p)
		if err != nil {
			return reflect.Value{}, err
		}
	}

	results := provider.value.Call(params)

	for _, r := range results {
		// Check error
		if isError(r.Type()) && !r.IsNil() {
			return reflect.Value{}, fmt.Errorf("%s call err: %+v", r, results)
		}
		if !isError(r.Type()) && !r.IsNil() {
			c.results[r.Type()] = r
		}
	}
	return c.results[param], nil
}

func (c *Container) Invoke(f reflect.Type) error {
	if f.Kind() == reflect.Func {
		return fmt.Errorf("f must be a funcion")
	}

	var err error
	params := make([]reflect.Value, f.NumIn())
	for i := range params {
		params[i], err = c.buildParam(f.In(i))
		if err != nil {
			return err
		}
	}

	reflect.ValueOf(f).Call(params)
	return nil
}
