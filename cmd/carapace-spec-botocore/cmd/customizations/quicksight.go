package customizations

import (
	_ "embed"

	"github.com/carapace-sh/carapace-spec/pkg/command"
)

func init() {
	customizations["quicksight.start-asset-bundle-import-job"] = func(cmd *command.Command) error {
		cmd.AddFlag(command.Flag{
			Longhand:    "asset-bundle-import-source-bytes",
			Description: "The content of the asset bundle to be uploaded.",
			Value:       true,
		})
		return nil
	}
}
