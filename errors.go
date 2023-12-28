package mazego

type mazeError string

func (me mazeError) Error() string {
	return string(me)
}

func (me mazeError) Is(target error) bool {
	return me == target
}
