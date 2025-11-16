package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_createContainerServiceDeploymentCmd = &cobra.Command{
	Use:   "create-container-service-deployment",
	Short: "Creates a deployment for your Amazon Lightsail container service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_createContainerServiceDeploymentCmd).Standalone()

	lightsail_createContainerServiceDeploymentCmd.Flags().String("containers", "", "An object that describes the settings of the containers that will be launched on the container service.")
	lightsail_createContainerServiceDeploymentCmd.Flags().String("public-endpoint", "", "An object that describes the settings of the public endpoint for the container service.")
	lightsail_createContainerServiceDeploymentCmd.Flags().String("service-name", "", "The name of the container service for which to create the deployment.")
	lightsail_createContainerServiceDeploymentCmd.MarkFlagRequired("service-name")
	lightsailCmd.AddCommand(lightsail_createContainerServiceDeploymentCmd)
}
