package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_createContainerServiceCmd = &cobra.Command{
	Use:   "create-container-service",
	Short: "Creates an Amazon Lightsail container service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_createContainerServiceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_createContainerServiceCmd).Standalone()

		lightsail_createContainerServiceCmd.Flags().String("deployment", "", "An object that describes a deployment for the container service.")
		lightsail_createContainerServiceCmd.Flags().String("power", "", "The power specification for the container service.")
		lightsail_createContainerServiceCmd.Flags().String("private-registry-access", "", "An object to describe the configuration for the container service to access private container image repositories, such as Amazon Elastic Container Registry (Amazon ECR) private repositories.")
		lightsail_createContainerServiceCmd.Flags().String("public-domain-names", "", "The public domain names to use with the container service, such as `example.com` and `www.example.com`.")
		lightsail_createContainerServiceCmd.Flags().String("scale", "", "The scale specification for the container service.")
		lightsail_createContainerServiceCmd.Flags().String("service-name", "", "The name for the container service.")
		lightsail_createContainerServiceCmd.Flags().String("tags", "", "The tag keys and optional values to add to the container service during create.")
		lightsail_createContainerServiceCmd.MarkFlagRequired("power")
		lightsail_createContainerServiceCmd.MarkFlagRequired("scale")
		lightsail_createContainerServiceCmd.MarkFlagRequired("service-name")
	})
	lightsailCmd.AddCommand(lightsail_createContainerServiceCmd)
}
