package dbclient

import (
	"context"
	"log"
	"time"

	"github.com/scardozos/add-weeks-tool/cmd/types"
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

func (s *LocalClient) GetStaticWeeks() []types.Date {
	c := s.db.Client
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := c.GetStaticWeeks(ctx, &weeks_pb.GetStaticWeeksRequest{})
	if err != nil {
		log.Print(err)
		return nil
	}
	ret := make([]types.Date, len(res.StaticWeeks))
	for idx, val := range res.StaticWeeks {
		ret[idx] = ToTypeDate(*val)
	}
	return ret
}

func (s *LocalClient) SetStaticWeek()   {}
func (s *LocalClient) UnsetStaticWeek() {}

func ToTypeDate(d weeks_pb.Date) types.Date {
	return types.Date{
		Day:   d.Day,
		Month: d.Month,
		Year:  d.Year,
	}
}
