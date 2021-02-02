package main

const write = "write"
const sleep = "sleep"
//CountdownOperationsSpy implements Sleeper, used to mock.
type CountdownOperationsSpy struct {
	Calls []string
}

//Sleep just do count
func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return 0,nil
}