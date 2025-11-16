package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supplychain_deleteDataIntegrationFlowCmd = &cobra.Command{
	Use:   "delete-data-integration-flow",
	Short: "Enable you to programmatically delete an existing data pipeline for the provided Amazon Web Services Supply Chain instance and DataIntegrationFlow name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supplychain_deleteDataIntegrationFlowCmd).Standalone()

	supplychain_deleteDataIntegrationFlowCmd.Flags().String("instance-id", "", "The Amazon Web Services Supply Chain instance identifier.")
	supplychain_deleteDataIntegrationFlowCmd.Flags().String("name", "", "The name of the DataIntegrationFlow to be deleted.")
	supplychain_deleteDataIntegrationFlowCmd.MarkFlagRequired("instance-id")
	supplychain_deleteDataIntegrationFlowCmd.MarkFlagRequired("name")
	supplychainCmd.AddCommand(supplychain_deleteDataIntegrationFlowCmd)
}
