package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_updateContainerServiceCmd = &cobra.Command{
	Use:   "update-container-service",
	Short: "Updates the configuration of your Amazon Lightsail container service, such as its power, scale, and public domain names.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_updateContainerServiceCmd).Standalone()

	lightsail_updateContainerServiceCmd.Flags().String("is-disabled", "", "A Boolean value to indicate whether the container service is disabled.")
	lightsail_updateContainerServiceCmd.Flags().String("power", "", "The power for the container service.")
	lightsail_updateContainerServiceCmd.Flags().String("private-registry-access", "", "An object to describe the configuration for the container service to access private container image repositories, such as Amazon Elastic Container Registry (Amazon ECR) private repositories.")
	lightsail_updateContainerServiceCmd.Flags().String("public-domain-names", "", "The public domain names to use with the container service, such as `example.com` and `www.example.com`.")
	lightsail_updateContainerServiceCmd.Flags().String("scale", "", "The scale for the container service.")
	lightsail_updateContainerServiceCmd.Flags().String("service-name", "", "The name of the container service to update.")
	lightsail_updateContainerServiceCmd.MarkFlagRequired("service-name")
	lightsailCmd.AddCommand(lightsail_updateContainerServiceCmd)
}
