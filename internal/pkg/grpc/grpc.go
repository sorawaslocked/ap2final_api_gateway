package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"time"
)

type (
	Config struct {
		Client Client
	}

	Client struct {
		MovieServiceURL  string        `yaml:"movieServiceURL" env:"GRPC_MOVIE_SERVICE_URL" env-required:"true"`
		MaxReceiveSizeMb int           `yaml:"maxReceiveSizeMb" env:"GRPC_MAX_RECEIVE_SIZE_MB" env-default:"4"`
		TimeKeepAlive    time.Duration `yaml:"timeKeepAlive" env:"GRPC_TIME_KEEP_ALIVE" env-default:"1m"`
		Timeout          time.Duration `yaml:"timeout" env:"GRPC_TIMEOUT" env-default:"10s"`
	}
)

func Connect(target string, clientCfg Client) (*grpc.ClientConn, error) {
	keepAliveParams := keepalive.ClientParameters{
		Time:                clientCfg.TimeKeepAlive,
		Timeout:             clientCfg.Timeout,
		PermitWithoutStream: true,
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepAliveParams),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(clientCfg.MaxReceiveSizeMb)),
	}

	conn, err := grpc.NewClient(target, opts...)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
