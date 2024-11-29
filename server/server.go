package server

import (
	pb "boost/im_repo_boost"
	"boost/internal/handlers"
	"boost/pkg/config"
	"boost/pkg/database"
	"context"
	"net"
	"sync"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Server struct {
	GrpcServer *grpc.Server
	wg         sync.WaitGroup
	lis        net.Listener
	logger     *logrus.Entry
	Cfg        *config.Config
}

func NewServer(ctx context.Context, logger *logrus.Entry, cfg *config.Config) *Server {
	lis, err := net.Listen("tcp", cfg.GrpcRepo.Port)

	if err != nil {
		logger.Fatalf("Failed to listen: %v", err)
	}
	s := &Server{
		GrpcServer: grpc.NewServer(),
		lis:        lis,
		logger:     logger,
		Cfg:        cfg,
	}

	register(ctx, s)

	return s
}

func (s *Server) GracefulShutdown(ctx context.Context) {
	s.logger.Info("Shutting down gRPC server gracefully...")
	s.GrpcServer.GracefulStop()
	done := make(chan struct{})

	go func() {
		defer close(done)
		s.wg.Wait() // Wait for all ongoing operations
	}()

	select {
	case <-done:
		s.logger.Info("All operations completed; gRPC server has shut down.")
	case <-ctx.Done():
		s.logger.Warn("force shutdown; exiting with ongoing operations.")
	}
}

func register(ctx context.Context, s *Server) {

	pool, err := database.NewClient(ctx, &database.Options{
		Host:          s.Cfg.Storage.Psql.Host,
		Port:          s.Cfg.Storage.Psql.Port,
		Database:      s.Cfg.Storage.Psql.Database,
		Username:      s.Cfg.Storage.Psql.Username,
		Password:      s.Cfg.Storage.Psql.Password,
		PgPoolMaxConn: s.Cfg.Storage.Psql.PgPoolMaxConn,
	})

	if err != nil {
		s.logger.Error(err)
	}

	handler := handlers.NewHandler(s.logger, pool)
	pb.RegisterBoostServiceServer(s.GrpcServer, handler)
}

func (s *Server) Serve() {
	go func() {
		s.logger.Info("Starting gRPC server...")

		if err := s.GrpcServer.Serve(s.lis); err != nil {
			s.logger.Fatalf("Failed to serve: %v", err)
		}
	}()
	s.logger.Info("gRPC server started on port ", s.Cfg.GrpcRepo.Port)

}

func (s *Server) Close() {
	s.logger.Info("Shutting down server...")
	s.GrpcServer.GracefulStop()

	s.logger.Info("Waiting for ongoing tasks to complete...")
	s.wg.Wait()

	s.logger.Info("Server stopped gracefully")
}
