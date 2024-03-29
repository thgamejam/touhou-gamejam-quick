package server

import (
    gameV1 "game/api/game/v1"
    userV1 "game/api/user/v1"
    "game/internal/conf"
    "game/internal/service"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/go-kratos/kratos/v2/middleware/recovery"
    "github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server,
    gameService *service.GameService, userService *service.UserService, logger log.Logger) *http.Server {
    
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
    gameV1.RegisterGameHTTPServer(srv, gameService)
    userV1.RegisterUserHTTPServer(srv, userService)
    
    return srv
}
