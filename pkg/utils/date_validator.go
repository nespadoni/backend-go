package utils

import (
	"fmt"
	"time"
)

// ValidateEventDates valida datas de início e fim de eventos
func ValidateEventDates(startDate, endDate time.Time) error {
	if endDate.Before(startDate) {
		return fmt.Errorf("data de término deve ser posterior à data de início")
	}

	if startDate.Before(time.Now().Truncate(24 * time.Hour)) {
		return fmt.Errorf("data de início não pode ser no passado")
	}

	return nil
}

// ValidateFutureDate valida se uma data está no futuro
func ValidateFutureDate(date time.Time) error {
	if date.Before(time.Now().Truncate(24 * time.Hour)) {
		return fmt.Errorf("data não pode estar no passado")
	}
	return nil
}

// ValidateDateSequence valida sequência de datas
func ValidateDateSequence(dates ...time.Time) error {
	for i := 1; i < len(dates); i++ {
		if dates[i].Before(dates[i-1]) {
			return fmt.Errorf("datas devem estar em ordem cronológica")
		}
	}
	return nil
}
