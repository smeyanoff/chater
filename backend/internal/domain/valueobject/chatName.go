package valueobject

import (
	"chater/internal/domain/validation"
	"database/sql/driver"
	"fmt"
)

// ChatName определяет пользовательский тип для имени чата
type ChatName struct {
	value string
}

// Реализация интерфейса driver.Valuer для сохранения в базу данных
func (cn ChatName) Value() (driver.Value, error) {
	return cn.value, nil
}

// Реализация интерфейса sql.Scanner для извлечения из базы данных
func (cn *ChatName) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("cannot scan %T into ChatName", value)
	}
	cn.value = str
	return nil
}

// Метод, чтобы получать строковое значение
func (cn ChatName) String() string {
	return cn.value
}

func NewChatName(name string) (ChatName, error) {
	err := validation.ValidateChatName(name)
	if err != nil {
		return ChatName{}, err
	}
	return ChatName{value: name}, nil
}
