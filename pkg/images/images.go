package images

import (
	"errors"
	"log"
	"os"
	"path"

	"github.com/cjdenio/dockerenv/graph/model"
	"gopkg.in/yaml.v2"
)

var LoadedImages Images

type Images struct {
	Images []*model.Image `json:"images"`
}

func (i *Images) FindOne(name string) (*model.Image, error) {
	for _, v := range i.Images {
		if v.Name == name {
			return v, nil
		}
	}

	return nil, errors.New("image not found")
}

func LoadImages(dir string) error {
	dirents, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	var images []*model.Image

	for _, v := range dirents {
		if !v.IsDir() {
			raw, err := os.ReadFile(path.Join(dir, v.Name()))
			if err != nil {
				log.Printf("WARN: error reading image `%s`\n", v.Name())
				continue
			}

			var image struct {
				Name      string
				Url       string
				Variables map[string]model.Variable
			}

			err = yaml.Unmarshal(raw, &image)
			if err != nil {
				log.Printf("WARN: error reading image `%s`: %s\n", v.Name(), err.Error())
			}

			var variables []*model.Variable

			for key, value := range image.Variables {
				variables = append(variables, &model.Variable{
					Name:        key,
					Description: value.Description,
					Default:     value.Default,
					Required:    value.Required,
					Uncommon:    value.Uncommon,
				})
			}

			images = append(images, &model.Image{
				Name:      image.Name,
				URL:       image.Url,
				Variables: variables,
			})
		}
	}

	LoadedImages = Images{
		images,
	}

	return nil
}
