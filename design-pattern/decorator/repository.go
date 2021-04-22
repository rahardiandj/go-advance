package decorator

import (
	"errors"
	"time"
)

type Args map[string]string

type Data map[string]string

type Fetcher interface {
	Fetch(Args) (Data, error)
}

type Repository struct {
}

func (r *Repository) Fetch(args Args) (Data, error) {

	if args == nil {
		return nil, errors.New("no argument provided")
	}

	return Data{
		"username" : "rah",
		"email" : "rah@mail.com"
	}, nil

}


type Retrier struct {
	RetryCount int
	WaitInterval time.Duration
	Fetcher Fetcher
}

func (r *Retrier) Fetch(args Args) (Data, error) {
	for retry := 1; retry <= r.RetryCount; retry++ {
		if data, err := Fetcher.Fetch(); err == nil {
			return data, nil
		} else if retry == r.RetryCount {
			return Data{}, nil
		}

		time.Sleep(r.WaitInterval)

	}
	return Data{}, nil
}
