package server

import (
	context "context"
	"github.com/rzeAkbari/observabilityGo/server/grpc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"io"
	"net"
)

type routeGuideServer struct {
	observabilityGo.UnimplementedRouteGuideServer
	savedFeatures []observabilityGo.Feature
	routeNotes    map[string][]*observabilityGo.RouteNote
}

func (s *routeGuideServer) GetFeature(ctx context.Context, point *observabilityGo.Point) (*observabilityGo.Feature, error) {
	for _, f := range s.savedFeatures {
		if proto.Equal(f.Point, point) {
			return &f, nil
		}
	}
	return &observabilityGo.Feature{Point: point}, nil
}

func (s *routeGuideServer) ListFeatures(rectangle *observabilityGo.Rectangle, stream observabilityGo.RouteGuide_ListFeaturesServer) error {
	for _, f := range s.savedFeatures {
		if f.Point.Latitude < rectangle.Lo.Latitude &&
			f.Point.Longitude > rectangle.Hi.Longitude {
			if err := stream.Send(&f); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *routeGuideServer) RecordRouts(stream observabilityGo.RouteGuide_RecordRoutsServer) error {
	rs := observabilityGo.RouteSummary{}
	for {
		point, err := stream.Recv()
		if err == io.EOF {
			stream.SendAndClose(&rs)
		}
		if err != nil {
			return err
		}
		rs.PointCount++
		for _, f := range s.savedFeatures {
			if proto.Equal(f.Point, point) {
				rs.FeatureCount++
			}
		}
	}
}

func (s *routeGuideServer) RouteChat(stream observabilityGo.RouteGuide_RouteChatServer) error {
	//TODO implement me
	panic("implement me")
}

func (s *routeGuideServer) mustEmbedUnimplementedRouteGuideServer() {
	//TODO implement me
	panic("implement me")
}

func newServer() *routeGuideServer {
	return &routeGuideServer{
		routeNotes:    map[string][]*observabilityGo.RouteNote{},
		savedFeatures: nil,
	}
}

func GrpcServe() {
	lis, _ := net.Listen("tcp", "localhost:9090")
	grpcServer := grpc.NewServer()
	observabilityGo.RegisterRouteGuideServer(grpcServer, newServer())
	grpcServer.Serve(lis)

}
