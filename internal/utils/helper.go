package utils

import (
    "regexp"
    "strings"
	"path/filepath"
)

// NormalizeName keeps spaces but removes unsafe characters and lowercases the string
func NormalizeName(name string) string {
    normalized := strings.ToLower(name)

    // Allow letters, numbers, spaces, dots, underscores, and hyphens
    re := regexp.MustCompile(`[^a-z0-9 ._-]+`)
    normalized = re.ReplaceAllString(normalized, "")

    normalized = strings.TrimSpace(normalized)

    return normalized
}

func NormalizeFilename(name string) string {
    normalized := strings.ToLower(name)

    // Allow letters, numbers, spaces, dots, underscores, and hyphens
    re := regexp.MustCompile(`[^a-z0-9 ._-]+`)
    normalized = re.ReplaceAllString(normalized, "")

    normalized = strings.TrimSpace(normalized)

    return normalized
}




func BuildAlbumPath(baseDir, artistName, albumName string) string {
    artistFolder := NormalizeName(artistName)
    albumFolder := NormalizeName(albumName)

    return filepath.Join(baseDir, artistFolder, albumFolder)
}