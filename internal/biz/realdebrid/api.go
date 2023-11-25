package realdebrid

/*
 * Request and Response bodies
 */

type LinkCheckRequest struct {
	Link string `structs:"link"`
}

type LinkCheckResponse struct {
	Host      string `json:"host"`
	Link      string `json:"link"`
	Filename  string `json:"filename"`
	FileSize  int    `json:"filesize"`
	Supported int    `json:"supported"`
}

type FolderUnrestrictRequest struct {
	Link string `structs:"link"`
}

type FolderUnrestrictResponse = []string

type LinkUnrestrictRequest struct {
	Link string `structs:"link"`
}

type LinkUnrestrictResponse struct {
	Id         string `json:"id"`
	Filename   string `json:"filename"`
	MimeType   string `json:"mimeType"`
	FileSize   int    `json:"filesize"`
	Link       string `json:"link"`
	Host       string `json:"host"`
	Chunks     int    `json:"chunks"`
	Supported  int    `json:"supported"`
	Crc        int    `json:"crc"`
	Download   string `json:"download"`
	Streamable int    `json:"streamable"`
}
