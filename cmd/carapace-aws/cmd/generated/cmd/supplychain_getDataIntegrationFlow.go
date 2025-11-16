package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supplychain_getDataIntegrationFlowCmd = &cobra.Command{
	Use:   "get-data-integration-flow",
	Short: "Enables you to programmatically view a specific data pipeline for the provided Amazon Web Services Supply Chain instance and DataIntegrationFlow name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supplychain_getDataIntegrationFlowCmd).Standalone()

	supplychain_getDataIntegrationFlowCmd.Flags().String("instance-id", "", "The Amazon Web Services Supply Chain instance identifier.")
	supplychain_getDataIntegrationFlowCmd.Flags().String("name", "", "The name of the DataIntegrationFlow created.")
	supplychain_getDataIntegrationFlowCmd.MarkFlagRequired("instance-id")
	supplychain_getDataIntegrationFlowCmd.MarkFlagRequired("name")
	supplychainCmd.AddCommand(supplychain_getDataIntegrationFlowCmd)
}
