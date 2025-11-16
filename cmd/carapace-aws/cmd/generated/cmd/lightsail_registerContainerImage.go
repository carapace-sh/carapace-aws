package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_registerContainerImageCmd = &cobra.Command{
	Use:   "register-container-image",
	Short: "Registers a container image to your Amazon Lightsail container service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_registerContainerImageCmd).Standalone()

	lightsail_registerContainerImageCmd.Flags().String("digest", "", "The digest of the container image to be registered.")
	lightsail_registerContainerImageCmd.Flags().String("label", "", "The label for the container image when it's registered to the container service.")
	lightsail_registerContainerImageCmd.Flags().String("service-name", "", "The name of the container service for which to register a container image.")
	lightsail_registerContainerImageCmd.MarkFlagRequired("digest")
	lightsail_registerContainerImageCmd.MarkFlagRequired("label")
	lightsail_registerContainerImageCmd.MarkFlagRequired("service-name")
	lightsailCmd.AddCommand(lightsail_registerContainerImageCmd)
}
