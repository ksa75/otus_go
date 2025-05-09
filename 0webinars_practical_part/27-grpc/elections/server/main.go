package main

import (
	"context"
	"github.com/OtusGolang/webinars_practical_part/27-grpc/elections/validate"
	"google.golang.org/grpc/metadata"
	"log"
	"net"

	"github.com/OtusGolang/webinars_practical_part/27-grpc/elections/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	pb.UnimplementedElectionsServer
}

func (s *Service) SubmitVote(ctx context.Context, req *pb.SubmitVoteRequest) (*pb.SubmitVoteResponse, error) {
	requestId := ""
	if ctx != nil {
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			ids := md.Get("request_id")
			if len(ids) > 0 {
				requestId = ids[0]
			}
		}
	vote := req.Vote
	if vote == nil {
		return nil, status.Error(codes.InvalidArgument, "vote is not specified")
	}

	}
	log.Printf("new vote receive (passport=%s, candidate_id=%d, time=%v, request_id: %v)",
		req.Vote.Passport, req.Vote.CandidateId, req.Vote.Time.AsTime(), requestId)

	log.Printf("vote accepted")
	return &pb.SubmitVoteResponse{}, nil
}

func main() {
	lsn, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			validate.UnaryServerRequestValidatorInterceptor(validate.ValidateReq),
			),
		)
	pb.RegisterElectionsServer(server, new(Service))

	log.Printf("starting server on %s", lsn.Addr().String())
	if err := server.Serve(lsn); err != nil {
		log.Fatal(err)
	}
}
