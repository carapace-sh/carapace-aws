package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appflow_createFlowCmd = &cobra.Command{
	Use:   "create-flow",
	Short: "Enables your application to create a new flow using Amazon AppFlow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appflow_createFlowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appflow_createFlowCmd).Standalone()

		appflow_createFlowCmd.Flags().String("client-token", "", "The `clientToken` parameter is an idempotency token.")
		appflow_createFlowCmd.Flags().String("description", "", "A description of the flow you want to create.")
		appflow_createFlowCmd.Flags().String("destination-flow-config-list", "", "The configuration that controls how Amazon AppFlow places data in the destination connector.")
		appflow_createFlowCmd.Flags().String("flow-name", "", "The specified name of the flow.")
		appflow_createFlowCmd.Flags().String("kms-arn", "", "The ARN (Amazon Resource Name) of the Key Management Service (KMS) key you provide for encryption.")
		appflow_createFlowCmd.Flags().String("metadata-catalog-config", "", "Specifies the configuration that Amazon AppFlow uses when it catalogs the data that's transferred by the associated flow.")
		appflow_createFlowCmd.Flags().String("source-flow-config", "", "The configuration that controls how Amazon AppFlow retrieves data from the source connector.")
		appflow_createFlowCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for your flow.")
		appflow_createFlowCmd.Flags().String("tasks", "", "A list of tasks that Amazon AppFlow performs while transferring the data in the flow run.")
		appflow_createFlowCmd.Flags().String("trigger-config", "", "The trigger settings that determine how and when the flow runs.")
		appflow_createFlowCmd.MarkFlagRequired("destination-flow-config-list")
		appflow_createFlowCmd.MarkFlagRequired("flow-name")
		appflow_createFlowCmd.MarkFlagRequired("source-flow-config")
		appflow_createFlowCmd.MarkFlagRequired("tasks")
		appflow_createFlowCmd.MarkFlagRequired("trigger-config")
	})
	appflowCmd.AddCommand(appflow_createFlowCmd)
}
