package actions

func (as *ActionSuite) Test_HealthCheck() {
	res := as.JSON("/healthz").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), `"status":"ok"`)
}
