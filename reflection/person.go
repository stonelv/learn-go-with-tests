package reflection

//Person is a struct shows person info
type Person struct {
	Name    string
	Profile Profile
}

//Profile is a struct shows person's profile info
type Profile struct {
	Age  int
	City string
}
