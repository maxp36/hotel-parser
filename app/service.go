package app

// Parser represent the parser's service
type Parser interface {
	JSONParser
	CSVParser
	XMLParser
}

// JSONParser represent the JSON parser's service
type JSONParser interface {
	ParseJSON(data []byte) error
}

// CSVParser represent the CSV parser's service
type CSVParser interface {
	ParseCSV(columns, data []string) error
}

// XMLParser represent the XML parser's service
type XMLParser interface {
	ParseXML(data []byte) error
}
