package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appflow_startFlowCmd = &cobra.Command{
	Use:   "start-flow",
	Short: "Activates an existing flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appflow_startFlowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appflow_startFlowCmd).Standalone()

		appflow_startFlowCmd.Flags().String("client-token", "", "The `clientToken` parameter is an idempotency token.")
		appflow_startFlowCmd.Flags().String("flow-name", "", "The specified name of the flow.")
		appflow_startFlowCmd.MarkFlagRequired("flow-name")
	})
	appflowCmd.AddCommand(appflow_startFlowCmd)
}
