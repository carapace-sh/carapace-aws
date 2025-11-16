package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appflow_cancelFlowExecutionsCmd = &cobra.Command{
	Use:   "cancel-flow-executions",
	Short: "Cancels active runs for a flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appflow_cancelFlowExecutionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appflow_cancelFlowExecutionsCmd).Standalone()

		appflow_cancelFlowExecutionsCmd.Flags().String("execution-ids", "", "The ID of each active run to cancel.")
		appflow_cancelFlowExecutionsCmd.Flags().String("flow-name", "", "The name of a flow with active runs that you want to cancel.")
		appflow_cancelFlowExecutionsCmd.MarkFlagRequired("flow-name")
	})
	appflowCmd.AddCommand(appflow_cancelFlowExecutionsCmd)
}
