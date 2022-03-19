package root

var (
	mp = map[string]string{
		"tinyint":    "int",
		"smallint":   "int",
		"mediumint":  "int",
		"int":        "int",
		"integer":    "int",
		"bigint":     "int",
		"float":      "float",
		"double":     "float",
		"decimal":    "string",
		"date":       "time.Time",
		"time":       "time.Time",
		"year":       "time.Time",
		"datetime":   "time.Time",
		"timestamp":  "int",
		"char":       "string",
		"varchar":    "string",
		"tinyblob":   "string",
		"tinytext":   "string",
		"blob":       "string",
		"text":       "string",
		"mediumblob": "string",
		"mediumtext": "string",
		"longblob":   "string",
		"longtext":   "string",
	}
)

func GetMaps() map[string]string {
	return mp
}
