package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supplychain_updateDataIntegrationFlowCmd = &cobra.Command{
	Use:   "update-data-integration-flow",
	Short: "Enables you to programmatically update an existing data pipeline to ingest data from the source systems such as, Amazon S3 buckets, to a predefined Amazon Web Services Supply Chain dataset (product, inbound\\_order) or a temporary dataset along with the data transformation query provided with the API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supplychain_updateDataIntegrationFlowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(supplychain_updateDataIntegrationFlowCmd).Standalone()

		supplychain_updateDataIntegrationFlowCmd.Flags().String("instance-id", "", "The Amazon Web Services Supply Chain instance identifier.")
		supplychain_updateDataIntegrationFlowCmd.Flags().String("name", "", "The name of the DataIntegrationFlow to be updated.")
		supplychain_updateDataIntegrationFlowCmd.Flags().String("sources", "", "The new source configurations for the DataIntegrationFlow.")
		supplychain_updateDataIntegrationFlowCmd.Flags().String("target", "", "The new target configurations for the DataIntegrationFlow.")
		supplychain_updateDataIntegrationFlowCmd.Flags().String("transformation", "", "The new transformation configurations for the DataIntegrationFlow.")
		supplychain_updateDataIntegrationFlowCmd.MarkFlagRequired("instance-id")
		supplychain_updateDataIntegrationFlowCmd.MarkFlagRequired("name")
	})
	supplychainCmd.AddCommand(supplychain_updateDataIntegrationFlowCmd)
}
