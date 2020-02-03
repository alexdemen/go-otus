package store

type ErrDateBusy struct {
}

func (e ErrDateBusy) Error() string {
	return "На заданный временной интервал уже существует событие."
}
