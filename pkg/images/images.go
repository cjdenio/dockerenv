package images

import (
	"errors"
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

type Variable struct {
	Description string `json:"description"`
	Required    bool   `json:"required"`
	Default     string `json:"default,omitempty"`
	Uncommon    bool   `json:"uncommon"`
}

type Image struct {
	Name      string              `json:"name"`
	Url       string              `json:"url"`
	Variables map[string]Variable `json:"variables"`
}

type Images struct {
	Images []Image `json:"images"`
}

func (i *Images) FindOne(name string) (Image, error) {
	for _, v := range i.Images {
		if v.Name == name {
			return v, nil
		}
	}

	return Image{}, errors.New("Image not found")
}

func LoadImages(dir string) (*Images, error) {
	dirents, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var images []Image

	for _, v := range dirents {
		if !v.IsDir() {
			raw, err := os.ReadFile(path.Join(dir, v.Name()))
			if err != nil {
				log.Println("WARN: error reading image `%s`", v.Name())
				continue
			}

			var image Image
			yaml.Unmarshal(raw, &image)

			images = append(images, image)
		}
	}

	return &Images{
		Images: images,
	}, nil
}
