package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supplychain_getDataIntegrationFlowExecutionCmd = &cobra.Command{
	Use:   "get-data-integration-flow-execution",
	Short: "Get the flow execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supplychain_getDataIntegrationFlowExecutionCmd).Standalone()

	supplychain_getDataIntegrationFlowExecutionCmd.Flags().String("execution-id", "", "The flow execution identifier.")
	supplychain_getDataIntegrationFlowExecutionCmd.Flags().String("flow-name", "", "The flow name.")
	supplychain_getDataIntegrationFlowExecutionCmd.Flags().String("instance-id", "", "The AWS Supply Chain instance identifier.")
	supplychain_getDataIntegrationFlowExecutionCmd.MarkFlagRequired("execution-id")
	supplychain_getDataIntegrationFlowExecutionCmd.MarkFlagRequired("flow-name")
	supplychain_getDataIntegrationFlowExecutionCmd.MarkFlagRequired("instance-id")
	supplychainCmd.AddCommand(supplychain_getDataIntegrationFlowExecutionCmd)
}
