package server

import (
    downloadV1 "download/api/download/v1"
    helloV1 "download/api/helloworld/v1"
    "download/internal/conf"
    "download/internal/service"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/go-kratos/kratos/v2/middleware/recovery"
    "github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server,
    greeter *service.GreeterService,
    download *service.DownloadService,
    logger log.Logger) *http.Server {

    var opts = []http.ServerOption{
        http.Middleware(
            recovery.Recovery(),
        ),
    }
    if c.Http.Network != "" {
        opts = append(opts, http.Network(c.Http.Network))
    }
    if c.Http.Addr != "" {
        opts = append(opts, http.Address(c.Http.Addr))
    }
    if c.Http.Timeout != nil {
        opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
    }
    srv := http.NewServer(opts...)
    helloV1.RegisterGreeterHTTPServer(srv, greeter)
    downloadV1.RegisterDownloadHTTPServer(srv, download)

    return srv
}