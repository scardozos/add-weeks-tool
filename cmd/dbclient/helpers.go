package dbclient

import weeks_pb "github.com/scardozos/esplai-weeks-db/api/weeksdb"

func ToTypeDate(d weeks_pb.Date) Date {
	return Date{
		Day:   d.Day,
		Month: d.Month,
		Year:  d.Year,
	}
}

func ToGrpcDate(d Date) *weeks_pb.Date {
	return &weeks_pb.Date{
		Day:   d.Day,
		Month: d.Month,
		Year:  d.Year,
	}
}
