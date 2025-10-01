package utils

import (
	"path/filepath"
	"regexp"
	"strings"
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

func NormalizeStringForYT(name string) string {
	name = strings.ToLower(name)

	// Allow only alphanumerics, spaces, dots, underscores, and hyphens
	clean := regexp.MustCompile(`[^a-z0-9 ._-]+`)
	name = clean.ReplaceAllString(name, "")

	// Collapse multiple spaces
	space := regexp.MustCompile(`\s+`)
	name = space.ReplaceAllString(name, " ")

	name = strings.TrimSpace(name)

	// Replace spaces with +
	name = strings.ReplaceAll(name, " ", "+")

	return name
}

func BuildAlbumPath(baseDir, artistName, albumName string) string {
	artistFolder := NormalizeName(artistName)
	albumFolder := NormalizeName(albumName)

	return filepath.Join(baseDir, artistFolder, albumFolder)
}

func BuildArtistPath(baseDir, artistName string) string {
	artistFolder := NormalizeName(artistName)
	return filepath.Join(baseDir, artistFolder)
}
