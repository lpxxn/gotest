//go:generate mockgen -destination ./dao/u_mock.go -package mock_data -source user.go

package d1

type User interface {
	Name() string
	SetAge(age int) bool
	V(idx int, name string) (string, error)
}

// go generate ./...
