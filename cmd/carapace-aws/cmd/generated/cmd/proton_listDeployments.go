package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_listDeploymentsCmd = &cobra.Command{
	Use:   "list-deployments",
	Short: "List deployments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_listDeploymentsCmd).Standalone()

	proton_listDeploymentsCmd.Flags().String("component-name", "", "The name of a component for result list filtering.")
	proton_listDeploymentsCmd.Flags().String("environment-name", "", "The name of an environment for result list filtering.")
	proton_listDeploymentsCmd.Flags().String("max-results", "", "The maximum number of deployments to list.")
	proton_listDeploymentsCmd.Flags().String("next-token", "", "A token that indicates the location of the next deployment in the array of deployment, after the list of deployment that was previously requested.")
	proton_listDeploymentsCmd.Flags().String("service-instance-name", "", "The name of a service instance for result list filtering.")
	proton_listDeploymentsCmd.Flags().String("service-name", "", "The name of a service for result list filtering.")
	protonCmd.AddCommand(proton_listDeploymentsCmd)
}
