package objects

type (
	// User model info
	// @Description Информация по пользователю
	User struct {
		// ID пользователя
		ID int64 `json:"id" example:"12"`
		// Имя пользователя
		Name string `json:"name" example:"Иван"`
		// Фамилия пользователя
		Surname string `json:"surname" example:"Иванович"`
		// Отчество пользователя
		Patronymic string `json:"patronymic" example:"Иванов"`
	} //@name User
)
