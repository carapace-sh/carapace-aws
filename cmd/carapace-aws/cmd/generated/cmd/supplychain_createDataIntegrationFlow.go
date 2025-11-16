package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supplychain_createDataIntegrationFlowCmd = &cobra.Command{
	Use:   "create-data-integration-flow",
	Short: "Enables you to programmatically create a data pipeline to ingest data from source systems such as Amazon S3 buckets, to a predefined Amazon Web Services Supply Chain dataset (product, inbound\\_order) or a temporary dataset along with the data transformation query provided with the API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supplychain_createDataIntegrationFlowCmd).Standalone()

	supplychain_createDataIntegrationFlowCmd.Flags().String("instance-id", "", "The Amazon Web Services Supply Chain instance identifier.")
	supplychain_createDataIntegrationFlowCmd.Flags().String("name", "", "Name of the DataIntegrationFlow.")
	supplychain_createDataIntegrationFlowCmd.Flags().String("sources", "", "The source configurations for DataIntegrationFlow.")
	supplychain_createDataIntegrationFlowCmd.Flags().String("tags", "", "The tags of the DataIntegrationFlow to be created")
	supplychain_createDataIntegrationFlowCmd.Flags().String("target", "", "The target configurations for DataIntegrationFlow.")
	supplychain_createDataIntegrationFlowCmd.Flags().String("transformation", "", "The transformation configurations for DataIntegrationFlow.")
	supplychain_createDataIntegrationFlowCmd.MarkFlagRequired("instance-id")
	supplychain_createDataIntegrationFlowCmd.MarkFlagRequired("name")
	supplychain_createDataIntegrationFlowCmd.MarkFlagRequired("sources")
	supplychain_createDataIntegrationFlowCmd.MarkFlagRequired("target")
	supplychain_createDataIntegrationFlowCmd.MarkFlagRequired("transformation")
	supplychainCmd.AddCommand(supplychain_createDataIntegrationFlowCmd)
}
