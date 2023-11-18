package CLI

import (
	"fmt"
	"github.com/dyte-submissions/november-2023-hiring-rohanailoni/server/CLI/Models"
	"regexp"
	"strings"
	"time"
)

func PrepareGeneralQuery(level, message, resourceID, traceID, spanID, commit, parentResourceID Models.Valueset, timestamp, fromTimestam, toTimestamp string) (string, []interface{}) {

	//If level is defined then we are not required to check all;
	CheckDegubLog := true
	CheckErrorLog := true
	if level.HasValue() {
		//fmt.Println("Table print", strings.ToLower(level.RegularFlag))
		if strings.ToLower(level.RegularFlag) == "error" {
			CheckErrorLog = true
			CheckDegubLog = false
		}
		if strings.ToLower(level.RegularFlag) == "debug" {
			CheckDegubLog = true
			CheckErrorLog = false
		}

	}

	queryErrorTable := "SELECT * FROM ErrorLog WHERE 1=1"
	queryDebugTable := "SELECT * FROM DebugLog WHERE 1=1"
	ErrorargsList := []interface{}{}
	DebugargsList := []interface{}{}
	queryErrorTable, queryDebugTable, ErrorargsList, DebugargsList = createIntermediateTimeStampQuery("timestamp", fromTimestam, toTimestamp, timestamp, queryErrorTable, queryDebugTable, ErrorargsList, DebugargsList)
	queryErrorTable, queryDebugTable, ErrorargsList, DebugargsList = CreateIntermediateQuery("level", level, queryErrorTable, queryDebugTable, ErrorargsList, DebugargsList)
	queryErrorTable, queryDebugTable, ErrorargsList, DebugargsList = CreateIntermediateQuery("message", message, queryErrorTable, queryDebugTable, ErrorargsList, DebugargsList)
	queryErrorTable, queryDebugTable, ErrorargsList, DebugargsList = CreateIntermediateQuery("resourceID", resourceID, queryErrorTable, queryDebugTable, ErrorargsList, DebugargsList)
	queryErrorTable, queryDebugTable, ErrorargsList, DebugargsList = CreateIntermediateQuery("traceId", traceID, queryErrorTable, queryDebugTable, ErrorargsList, DebugargsList)
	queryErrorTable, queryDebugTable, ErrorargsList, DebugargsList = CreateIntermediateQuery("spanId", spanID, queryErrorTable, queryDebugTable, ErrorargsList, DebugargsList)
	queryErrorTable, queryDebugTable, ErrorargsList, DebugargsList = CreateIntermediateQuery("commit", commit, queryErrorTable, queryDebugTable, ErrorargsList, DebugargsList)
	queryErrorTable, queryDebugTable, ErrorargsList, DebugargsList = CreateIntermediateQuery("metadata_parentResourceId", parentResourceID, queryErrorTable, queryDebugTable, ErrorargsList, DebugargsList)

	//fmt.Println(queryErrorTable, queryDebugTable, ErrorargsList, DebugargsList, CheckErrorLog, CheckDegubLog)
	if CheckErrorLog && !CheckDegubLog {
		return queryErrorTable, ErrorargsList
	} else if !CheckErrorLog && CheckDegubLog {
		return queryDebugTable, DebugargsList
	}
	finalQuery := queryErrorTable + " UNION " + queryDebugTable
	return finalQuery, append(ErrorargsList, DebugargsList...)
}

func containsWildcard(s string) bool {
	return regexp.MustCompile(`%`).MatchString(s)
}
func isValidRegex(pattern string) bool {
	_, err := regexp.Compile(pattern)
	return err == nil
}
func CreateIntermediateQuery(types string, set Models.Valueset, queryErrorTable, queryDebugTable string, ErrorargsList, DebugargsList []interface{}) (string, string, []interface{}, []interface{}) {
	if set.HasValue() {
		if set.HasRegex() {
			queryErrorTable += fmt.Sprintf(" AND %s REGEXP ?", types)
			queryDebugTable += fmt.Sprintf(" AND %s REGEXP ?", types)
			ErrorargsList = append(ErrorargsList, set.RegexFlag)
			DebugargsList = append(DebugargsList, set.RegexFlag)
		} else if set.HasRegular() {
			queryErrorTable += fmt.Sprintf(" AND %s = ?", types)
			queryDebugTable += fmt.Sprintf(" AND %s = ?", types)
			ErrorargsList = append(ErrorargsList, set.RegularFlag)
			DebugargsList = append(DebugargsList, set.RegularFlag)
		} else if set.HasWildcard() {
			queryErrorTable += fmt.Sprintf(" AND %s LIKE ?", types)
			queryDebugTable += fmt.Sprintf(" AND %s LIKE ?", types)
			ErrorargsList = append(ErrorargsList, set.WildcardFlag)
			DebugargsList = append(DebugargsList, set.WildcardFlag)
		}
	}

	return queryErrorTable, queryDebugTable, ErrorargsList, DebugargsList
}
func createIntermediateTimeStampQuery(types string, from, to, timestamp, queryErrorTable, queryDebugTable string, ErrorargsList, DebugargsList []interface{}) (string, string, []interface{}, []interface{}) {
	if timestamp == "" && from == "" && to == "" {
		return queryErrorTable, queryDebugTable, ErrorargsList, DebugargsList
	}
	if from != "" && to == "" {
		//we are considering exact user time as to time
		to = time.Now().Format(time.RFC3339)
		queryErrorTable += fmt.Sprintf(" AND %s BETWEEN ? AND ?", types)
		queryDebugTable += fmt.Sprintf(" WHERE %s BETWEEN ? AND ?", types)
		ErrorargsList = append(ErrorargsList, from, to)
		DebugargsList = append(DebugargsList, from, to)
	} else if timestamp != "" {
		queryErrorTable += fmt.Sprintf(" AND %s = ?", types)
		queryDebugTable += fmt.Sprintf(" WHERE %s = ?", types)
		ErrorargsList = append(ErrorargsList, timestamp)
		DebugargsList = append(DebugargsList, timestamp)
	}
	return queryErrorTable, queryDebugTable, ErrorargsList, DebugargsList
}
