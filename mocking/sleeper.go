package main

//Sleeper is a interface used to separate the timer connection
type Sleeper interface{
	Sleep()
}