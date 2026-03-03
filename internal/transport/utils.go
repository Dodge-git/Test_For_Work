package transport

import (
	"strconv"
	"strings"
)

func extractID(path string) (uint, error) {
	parts := strings.Split(strings.Trim(path, "/"), "/")

	if len(parts) < 2 {
		return 0, strconv.ErrSyntax
	}

	idStr := parts[len(parts)-1]

	parsed, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint(parsed), nil
}
func extractDepartmentIDFromEmployeePath(path string) (uint, error) {
	
	parts := strings.Split(strings.Trim(path, "/"), "/")

	if len(parts) != 3 {
		return 0, strconv.ErrSyntax
	}

	idStr := parts[1]

	parsed, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint(parsed), nil
}