package dbclient

import (
	"context"
	"log"
	"time"

	weeks_pb "github.com/scardozos/esplai-weeks-db/api/weeksdb"
)

type LocalClient struct {
	db *DbClient
}

func NewLocalClient(host string, secure bool) *LocalClient {
	return &LocalClient{
		db: NewDbClient(host, secure),
	}
}

func (s *LocalClient) GetStaticWeeks() ([]Date, error) {
	//var wg sync.WaitGroup

	c := s.db.Client
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	now := time.Now()
	res, err := c.GetStaticWeeks(ctx, &weeks_pb.GetStaticWeeksRequest{})
	log.Printf("Got result in %s", time.Since(now))
	if err != nil {
		return nil, err
	}

	ret := make([]Date, len(res.StaticWeeks))

	for idx, val := range res.StaticWeeks {
		ret[idx] = ToTypeDate(*val)
	}
	return ret, nil
}

func (s *LocalClient) SetStaticWeek(d Date) (Date, error) {
	c := s.db.Client
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	staticWeek := ToGrpcDate(d)
	_, err := c.SetStaticWeek(ctx, &weeks_pb.SetStaticWeekRequest{
		StaticWeek: staticWeek,
	})
	if err != nil {
		return Date{}, err
	}
	return d, nil
}

func (s *LocalClient) UnsetStaticWeek(d Date) (Date, error) {
	c := s.db.Client
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	staticWeek := ToGrpcDate(d)
	_, err := c.UnsetStaticWeek(ctx, &weeks_pb.UnsetStaticWeekRequest{
		StaticWeek: staticWeek,
	})
	if err != nil {
		return Date{}, err
	}
	return d, nil
}
