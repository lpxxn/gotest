//go:generate mockgen -destination ./dao/u_mock.go -package mock_data -source=user.go

package d1

type User interface {
	V(idx int, name string) (string, error)
	Name() string
	SetAge(age int) bool
}

// go generate ./...
