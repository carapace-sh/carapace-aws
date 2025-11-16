package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getContainerServiceDeploymentsCmd = &cobra.Command{
	Use:   "get-container-service-deployments",
	Short: "Returns the deployments for your Amazon Lightsail container service\n\nA deployment specifies the settings, such as the ports and launch command, of containers that are deployed to your container service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getContainerServiceDeploymentsCmd).Standalone()

	lightsail_getContainerServiceDeploymentsCmd.Flags().String("service-name", "", "The name of the container service for which to return deployments.")
	lightsail_getContainerServiceDeploymentsCmd.MarkFlagRequired("service-name")
	lightsailCmd.AddCommand(lightsail_getContainerServiceDeploymentsCmd)
}
