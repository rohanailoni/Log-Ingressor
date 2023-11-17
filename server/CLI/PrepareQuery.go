package CLI

func PrepareGeneralQuery(level, message, resourceID, timestamp, traceID, spanID, commit, parentResourceID string) (string, []interface{}) {
	query := "SELECT * FROM ErrorLog WHERE 1=1"
	argsList := []interface{}{}
	if level != "" {
		query += " AND level = ?"
		argsList = append(argsList, level)
	}

	if message != "" {
		query += " AND message = ?"
		argsList = append(argsList, message)
	}

	if resourceID != "" {
		query += " AND resourceId = ?"
		argsList = append(argsList, resourceID)
	}

	if timestamp != "" {
		query += " AND timestamp = ?"
		argsList = append(argsList, timestamp)
	}

	if traceID != "" {
		query += " AND TraceId = ?"
		argsList = append(argsList, traceID)
	}

	if spanID != "" {
		query += " AND spanId = ?"
		argsList = append(argsList, spanID)
	}

	if commit != "" {
		query += " AND Commit = ?"
		argsList = append(argsList, commit)
	}

	if parentResourceID != "" {
		query += " AND metadata_parentResourceId = ?"
		argsList = append(argsList, parentResourceID)
	}
	return query, argsList
}
