package server

func (s *CronServer) addRoutes() {
	s.srv.POST("/cron", s.addCron)
}
