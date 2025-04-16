package utils

import "github.com/google/uuid"

func GenerateUUID() (uuid.UUID, error) {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return uuid.Nil, err
	}

	return newUUID, nil
}

func ParseUUID(s string) (uuid.UUID, error) {
	parseUUID, err := uuid.Parse(s)
	if err != nil {
		return uuid.Nil, err
	}

	return parseUUID, nil
}
