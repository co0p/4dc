package assets

import "encoding/base64"

// Icon returns a tiny 1x1 PNG used as a minimal tray icon.
func Icon() []byte {
	const b64 = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8Xw8AAn0B9oQe3/AAAAAASUVORK5CYII="
	data, _ := base64.StdEncoding.DecodeString(b64)
	return data
}
