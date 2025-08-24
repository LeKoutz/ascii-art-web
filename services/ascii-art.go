package services

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	charHeight = 8
	startASCII = 32
)

type AsciiArtWeb struct {
	banners map[string]map[rune][]string
}

func NewAsciiArtWeb() *AsciiArtWeb {
	return &AsciiArtWeb{
		banners: make(map[string]map[rune][]string),
	}
}

func (a *AsciiArtWeb) LoadBanners() error {
	files, err := os.ReadDir("./banners")
	if err != nil {
		return fmt.Errorf("failed reading banners directory: %w", err)
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".txt") {
			continue
		}

		name := strings.TrimSuffix(file.Name(), ".txt")
		bannerfile, err := os.Open("./banners/" + name + ".txt")
		if err != nil {
			return fmt.Errorf("failed loading banner: %w", err)
		}
		defer bannerfile.Close()

		banner := make(map[rune][]string)
		scanner := bufio.NewScanner(bannerfile)
		var lines []string
		currentRune := rune(startASCII)

		for scanner.Scan() {
			line := scanner.Text()
			
			if line == "" && len(lines) > 0 {
				// Save completed character
				if len(lines) == charHeight {
					banner[currentRune] = lines
					currentRune++
				}
				lines = []string{}
				continue
			}
			
			if line != "" {
				lines = append(lines, line)
			}
		}

		// Handle the last character
		if len(lines) == charHeight {
			banner[currentRune] = lines
		}

		if err := scanner.Err(); err != nil {
			return fmt.Errorf("failed loading file: %w", err)
		}

		a.banners[name] = banner
	}
	return nil
}

// Generate is now a method of AsciiArtWeb
func (a *AsciiArtWeb) Generate(text, bannerName string) (string, error) {
	banner, exists := a.banners[bannerName]
	if !exists {
		return "", fmt.Errorf("banner '%s' not loaded", bannerName)
	}

	var result strings.Builder
	linesToPrint := make([]string, charHeight)

	for _, char := range text {
		if char == '\n' {
			// Handle newline - flush current lines
			for i := 0; i < charHeight; i++ {
				if linesToPrint[i] != "" {
					result.WriteString(linesToPrint[i] + "\n")
					linesToPrint[i] = ""
				}
			}
			result.WriteString("\n")
			continue
		}

		charArt := banner[char]
		if charArt == nil {
			// Use space for unknown characters
			charArt = banner[' ']
		}

		for i := 0; i < charHeight; i++ {
			if i < len(charArt) {
				linesToPrint[i] += charArt[i]
			}
		}
	}

	// Flush any remaining lines
	for i := 0; i < charHeight; i++ {
		if linesToPrint[i] != "" {
			result.WriteString(linesToPrint[i] + "\n")
		}
	}

	return result.String(), nil
}

func (a *AsciiArtWeb) GetAvailableBanners() []string {
	var banners []string
	for name := range a.banners {
		banners = append(banners, name)
	}
	return banners
}