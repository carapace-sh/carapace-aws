package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_deleteContainerImageCmd = &cobra.Command{
	Use:   "delete-container-image",
	Short: "Deletes a container image that is registered to your Amazon Lightsail container service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_deleteContainerImageCmd).Standalone()

	lightsail_deleteContainerImageCmd.Flags().String("image", "", "The name of the container image to delete from the container service.")
	lightsail_deleteContainerImageCmd.Flags().String("service-name", "", "The name of the container service for which to delete a registered container image.")
	lightsail_deleteContainerImageCmd.MarkFlagRequired("image")
	lightsail_deleteContainerImageCmd.MarkFlagRequired("service-name")
	lightsailCmd.AddCommand(lightsail_deleteContainerImageCmd)
}
