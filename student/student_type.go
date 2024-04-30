package student

type Student struct {
	Name          string `db:"name"`
	ClassesLeft   int    `db:"classes_left"`
	CurrentLesson string `db:"current_lesson"`
	Platform      string `db:"platform"`
}
