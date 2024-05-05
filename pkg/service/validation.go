package service

import "regexp"

func IsValidNumber(number string) bool {
	regex := `^(\+7|7|8)?\s?\(?(9\d{2})\)?[-\s]?(\d{3})[-\s]?(\d{2})[-\s]?(\d{2})$`
	return regexp.MustCompile(regex).MatchString(number)
}

func IsValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(regex).MatchString(email)
}

func IsValidPassword(password string) bool {
	// Регулярное выражение для проверки пароля:
	// - Длина от 8 до 20 символов
	// - Содержит хотя бы одну заглавную букву
	// - Содержит хотя бы одну строчную букву
	// - Содержит хотя бы одну цифру
	// - Содержит хотя бы один специальный символ из списка !@#$%^&*()-_+=
	regex := `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*()-_+=]).{8,20}$`
	return regexp.MustCompile(regex).MatchString(password)
}
