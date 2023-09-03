package storage

type IStorage interface {
}

type IConn interface {
	Open() error
	Close() error
	isClose() bool
}
