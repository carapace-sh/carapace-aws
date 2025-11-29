package customizations

import (
	_ "embed"

	"github.com/carapace-sh/carapace-spec/pkg/command"
)

func init() {
	for _, operation := range []string{
		"rekognition.detect-custom-labels",
		"rekognition.detect-faces",
		"rekognition.detect-labels",
		"rekognition.detect-moderation-labels",
		"rekognition.detect-protective-equipment",
		"rekognition.detect-text",
		"rekognition.index-faces",
		"rekognition.recognize-celebrities",
		"rekognition.search-faces-by-image",
		"rekognition.search-users-by-image",
	} {
		customizations[operation] = func(cmd *command.Command) error {
			cmd.AddFlag(command.Flag{
				Longhand:    "image-bytes",
				Description: "The content of the image to be uploaded",
				Value:       true,
			})
			return nil
		}
	}

	customizations["rekognition.compare-faces"] = func(cmd *command.Command) error {
		cmd.AddFlag(command.Flag{
			Longhand:    "source-image-bytes",
			Description: "The content of the image to be uploaded",
			Value:       true,
		})
		cmd.AddFlag(command.Flag{
			Longhand:    "target-image-bytes",
			Description: "The content of the image to be uploaded",
			Value:       true,
		})
		return nil
	}
}
