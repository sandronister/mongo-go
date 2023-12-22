package catch

func Execute(err error) {
	if err != nil {
		panic(err)
	}
}
