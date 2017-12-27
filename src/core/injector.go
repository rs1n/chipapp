package core

import "github.com/facebookgo/inject"

func Inject(result interface{}, dependencies ...interface{}) error {
	objects := make([]*inject.Object, len(dependencies)+1)
	objects[0] = &inject.Object{Value: result}
	fillDependencyObjects(objects, 1, dependencies)

	var graph inject.Graph
	if err := graph.Provide(objects...); err != nil {
		return err
	}
	return graph.Populate()
}

func fillDependencyObjects(
	objects []*inject.Object, startIndex int, dependencies []interface{},
) {
	for i, dep := range dependencies {
		objects[i+startIndex] = &inject.Object{
			Value:    dep,
			Complete: true,
		}
	}
}
