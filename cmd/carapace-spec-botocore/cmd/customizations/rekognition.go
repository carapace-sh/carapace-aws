package customizations

import (
	_ "embed"

	"github.com/carapace-sh/carapace-spec/pkg/command"
)

func init() {
	for _, operation := range []string{
		"rekognition.search-users-by-image",
		"rekognition.detect-text",
		"rekognition.detect-protective-equipment",
		"rekognition.detect-moderation-labels",
		"rekognition.detect-faces",
		"rekognition.detect-labels",
		"rekognition.detect-custom-labels",
		"rekognition.search-faces-by-image",
		"rekognition.index-faces",
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
