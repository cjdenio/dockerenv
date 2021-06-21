package images

import (
	"errors"
	"log"
	"os"
	"path"
	"sort"
	"strings"

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

func (i *Images) Search(query string) []*model.Image {
	matched_images := []*model.Image{}

	if query == "" {
		return matched_images
	}

	for _, v := range i.Images {
		if strings.Contains(strings.ToLower(v.Name), strings.ToLower(query)) {
			matched_images = append(matched_images, v)
		}
	}

	sort.Slice(matched_images, func(i, j int) bool {
		if strings.EqualFold(matched_images[j].Name, query) {
			return false
		} else if strings.EqualFold(matched_images[i].Name, query) {
			return true
		}

		return matched_images[i].Name < matched_images[j].Name
	})

	return matched_images
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
				continue
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

			sort.Slice(variables, func(i, j int) bool {
				if !variables[j].Uncommon && variables[i].Uncommon {
					return false
				} else if variables[j].Uncommon && !variables[i].Uncommon {
					return true
				}

				return i < j
			})

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
