package login_test

func Login(uname string,passwd string) string {

	username := "sourabh@trellis"
	password := "sourabh123"

	if username == uname {
		if password == passwd {
			return "Login successfull."
		} else {
			return "Wrong credentials."
		}
	} else {
		return "Wrong credentials."
	}
	
}

