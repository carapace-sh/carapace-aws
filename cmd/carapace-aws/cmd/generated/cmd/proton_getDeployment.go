package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_getDeploymentCmd = &cobra.Command{
	Use:   "get-deployment",
	Short: "Get detailed data for a deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_getDeploymentCmd).Standalone()

	proton_getDeploymentCmd.Flags().String("component-name", "", "The name of a component that you want to get the detailed data for.")
	proton_getDeploymentCmd.Flags().String("environment-name", "", "The name of a environment that you want to get the detailed data for.")
	proton_getDeploymentCmd.Flags().String("id", "", "The ID of the deployment that you want to get the detailed data for.")
	proton_getDeploymentCmd.Flags().String("service-instance-name", "", "The name of the service instance associated with the given deployment ID.")
	proton_getDeploymentCmd.Flags().String("service-name", "", "The name of the service associated with the given deployment ID.")
	proton_getDeploymentCmd.MarkFlagRequired("id")
	protonCmd.AddCommand(proton_getDeploymentCmd)
}
