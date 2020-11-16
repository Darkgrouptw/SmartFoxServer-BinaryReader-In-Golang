package Common

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
