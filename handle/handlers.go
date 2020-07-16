package handle

// HandlerGen 集約
type HandlerGen struct {
	Smplhandler *SmplHandler
	Camhandler  *CamHandler
}

// InitHandlerGen 初期化
func InitHandlerGen(s *SmplHandler, c *CamHandler) *HandlerGen {
	return &HandlerGen{
		Smplhandler: s,
		Camhandler:  c,
	}
}
