package comms

import "strings"

func GetTargetTable(logLevel string) string {
	// Determine the target table based on the log level (sharding key)
	if strings.ToLower(logLevel) == LogError {
		return string(ErrorTable)
	} else if strings.ToLower(logLevel) == LogDebug {
		return string(DebugTable)
	}

	// Default to a generic table Like if it is not Error or Log
	return string(ErrorTable)

}
