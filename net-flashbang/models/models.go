package models

// Book persistence
var filePersistence bool = false

// EnableFilePersistence enables the file persistence
func EnableFilePersistence() {
	filePersistence = true
}

// DisableFilePersistence disables the file persistence
func DisableFilePersistence() {
	filePersistence = false
}
