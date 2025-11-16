package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appflow_stopFlowCmd = &cobra.Command{
	Use:   "stop-flow",
	Short: "Deactivates the existing flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appflow_stopFlowCmd).Standalone()

	appflow_stopFlowCmd.Flags().String("flow-name", "", "The specified name of the flow.")
	appflow_stopFlowCmd.MarkFlagRequired("flow-name")
	appflowCmd.AddCommand(appflow_stopFlowCmd)
}
