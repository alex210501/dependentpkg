package ds

import "github.com/alex210501/borbo.io/ds"

func New(name string, age int) ds.DataStucture {
	return ds.DataStucture{Name: name, Age: age}
}

func SetOlder(person *ds.DataStucture, addedAge int) {
	person.Age += addedAge
}

func SetYounger(person *ds.DataStucture, reducedAge int) {
	person.Age -= reducedAge
}
