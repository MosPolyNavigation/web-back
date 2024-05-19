package pkg

import "strconv"

func PickInt(query string) (int, error) {
	value, err := strconv.Atoi(query)

	if err != nil {
		return 0, err
	}

	return value, nil
}
