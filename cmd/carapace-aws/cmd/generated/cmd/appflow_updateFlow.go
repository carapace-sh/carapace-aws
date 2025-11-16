package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appflow_updateFlowCmd = &cobra.Command{
	Use:   "update-flow",
	Short: "Updates an existing flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appflow_updateFlowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appflow_updateFlowCmd).Standalone()

		appflow_updateFlowCmd.Flags().String("client-token", "", "The `clientToken` parameter is an idempotency token.")
		appflow_updateFlowCmd.Flags().String("description", "", "A description of the flow.")
		appflow_updateFlowCmd.Flags().String("destination-flow-config-list", "", "The configuration that controls how Amazon AppFlow transfers data to the destination connector.")
		appflow_updateFlowCmd.Flags().String("flow-name", "", "The specified name of the flow.")
		appflow_updateFlowCmd.Flags().String("metadata-catalog-config", "", "Specifies the configuration that Amazon AppFlow uses when it catalogs the data that's transferred by the associated flow.")
		appflow_updateFlowCmd.Flags().String("source-flow-config", "", "")
		appflow_updateFlowCmd.Flags().String("tasks", "", "A list of tasks that Amazon AppFlow performs while transferring the data in the flow run.")
		appflow_updateFlowCmd.Flags().String("trigger-config", "", "The trigger settings that determine how and when the flow runs.")
		appflow_updateFlowCmd.MarkFlagRequired("destination-flow-config-list")
		appflow_updateFlowCmd.MarkFlagRequired("flow-name")
		appflow_updateFlowCmd.MarkFlagRequired("source-flow-config")
		appflow_updateFlowCmd.MarkFlagRequired("tasks")
		appflow_updateFlowCmd.MarkFlagRequired("trigger-config")
	})
	appflowCmd.AddCommand(appflow_updateFlowCmd)
}
