package server

import (
	"context"

	pb "github.com/brotherlogic/photosbridge/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetAlbum(ctx context.Context, req *pb.GetAlbumRequest) (*pb.GetAlbumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "not done yet")
}
