package pkg

import "fmt"


func HandleError(ctx string, err error) string {
	if err != nil {
		return fmt.Sprintf("Получена ошибка в %s: %v", ctx, err)
	}
	return ""
}