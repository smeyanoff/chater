package valueobject

import (
	"chater/internal/domain/validation"
	"database/sql/driver"
	"fmt"
)

type GroupName struct {
	value string
}

// Реализация интерфейса driver.Valuer для сохранения в базу данных
func (gn GroupName) Value() (driver.Value, error) {
	return gn.value, nil
}

// Реализация интерфейса sql.Scanner для извлечения из базы данных
func (gn *GroupName) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("cannot scan %T into GroupName", value)
	}
	gn.value = str
	return nil
}

// Метод, чтобы получать строковое значение
func (gn GroupName) String() string {
	return gn.value
}

func NewGroupName(name string) (GroupName, error) {
	err := validation.ValidateGroupName(name)
	if err != nil {
		return GroupName{}, err
	}
	return GroupName{value: name}, nil
}
