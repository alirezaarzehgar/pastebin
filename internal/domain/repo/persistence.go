package repo

type Persistence interface {
	Connect() error
}
