package utils

// TypeByExtension returns the mime type for a given file extension
// source: https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types/Complete_list_of_MIME_types
func TypeByExtension(ext string) string {
	m, ok := typeMap[ext]

	if !ok {
		return "application/octet-stream" // default to octet-stream
	}

	return m
}

var typeMap = map[string]string{
	".aac":   "audio/aac",
	".abw":   "application/x-abiword",
	".arc":   "application/octet-stream",
	".avi":   "video/x-msvideo",
	".azw":   "application/vnd.amazon.ebook",
	".bin":   "application/octet-stream",
	".bz":    "application/x-bzip",
	".bz2":   "application/x-bzip2",
	".csh":   "application/x-csh",
	".css":   "text/css; charset=utf-8",
	".csv":   "text/csv",
	".doc":   "application/msword",
	".epub":  "application/epub+zip",
	".gif":   "image/gif",
	".htm":   "text/html; charset=utf-8",
	".html":  "text/html; charset=utf-8",
	".ico":   "image/x-icon",
	".ics":   "text/calendar",
	".jar":   "application/java-archive",
	".jpeg":  "image/jpeg",
	".jpg":   "image/jpeg",
	".js":    "application/x-javascript",
	".json":  "application/json",
	".mid":   "audio/midi",
	".midi":  "audio/midi",
	".mpeg":  "video/mpeg",
	".mpkg":  "application/vnd.apple.installer+xml",
	".odp":   "application/vnd.oasis.opendocument.presentation",
	".ods":   "application/vnd.oasis.opendocument.spreadsheet",
	".odt":   "application/vnd.oasis.opendocument.text",
	".oga":   "audio/ogg",
	".ogv":   "video/ogg",
	".ogx":   "application/ogg",
	".pdf":   "application/pdf",
	".ppt":   "application/vnd.ms-powerpoint",
	".rar":   "application/x-rar-compressed",
	".rtf":   "application/rtf",
	".sh":    "application/x-sh",
	".svg":   "image/svg+xml",
	".swf":   "application/x-shockwave-flash",
	".tar":   "application/x-tar",
	".tif":   "image/tiff",
	".tiff":  "image/tiff",
	".ttf":   "application/x-font-ttf",
	".txt":   "text/plain; charset=utf-8",
	".vsd":   "application/vnd.visio",
	".wav":   "audio/x-wav",
	".weba":  "audio/webm",
	".webm":  "video/webm",
	".webp":  "image/webp",
	".woff":  "application/x-font-woff",
	".xhtml": "application/xhtml+xml",
	".xls":   "application/vnd.ms-excel",
	".xml":   "application/xml; charset=utf-8",
	".xul":   "application/vnd.mozilla.xul+xml",
	".zip":   "application/zip",
	".3gp":   "audio/3gpp",
	".3g2":   "audio/3gpp2",
	".7z ":   "application/x-7z-compressed",
}
