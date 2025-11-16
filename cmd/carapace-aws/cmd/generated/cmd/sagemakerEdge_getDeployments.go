package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerEdge_getDeploymentsCmd = &cobra.Command{
	Use:   "get-deployments",
	Short: "Use to get the active deployments from a device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerEdge_getDeploymentsCmd).Standalone()

	sagemakerEdge_getDeploymentsCmd.Flags().String("device-fleet-name", "", "The name of the fleet that the device belongs to.")
	sagemakerEdge_getDeploymentsCmd.Flags().String("device-name", "", "The unique name of the device you want to get the configuration of active deployments from.")
	sagemakerEdge_getDeploymentsCmd.MarkFlagRequired("device-fleet-name")
	sagemakerEdge_getDeploymentsCmd.MarkFlagRequired("device-name")
	sagemakerEdgeCmd.AddCommand(sagemakerEdge_getDeploymentsCmd)
}
