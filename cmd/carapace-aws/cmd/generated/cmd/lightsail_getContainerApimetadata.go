package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getContainerApimetadataCmd = &cobra.Command{
	Use:   "get-container-apimetadata",
	Short: "Returns information about Amazon Lightsail containers, such as the current version of the Lightsail Control (lightsailctl) plugin.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getContainerApimetadataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_getContainerApimetadataCmd).Standalone()

	})
	lightsailCmd.AddCommand(lightsail_getContainerApimetadataCmd)
}
