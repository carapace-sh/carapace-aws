package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getContainerImagesCmd = &cobra.Command{
	Use:   "get-container-images",
	Short: "Returns the container images that are registered to your Amazon Lightsail container service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getContainerImagesCmd).Standalone()

	lightsail_getContainerImagesCmd.Flags().String("service-name", "", "The name of the container service for which to return registered container images.")
	lightsail_getContainerImagesCmd.MarkFlagRequired("service-name")
	lightsailCmd.AddCommand(lightsail_getContainerImagesCmd)
}
