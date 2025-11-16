package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appflow_describeFlowCmd = &cobra.Command{
	Use:   "describe-flow",
	Short: "Provides a description of the specified flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appflow_describeFlowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appflow_describeFlowCmd).Standalone()

		appflow_describeFlowCmd.Flags().String("flow-name", "", "The specified name of the flow.")
		appflow_describeFlowCmd.MarkFlagRequired("flow-name")
	})
	appflowCmd.AddCommand(appflow_describeFlowCmd)
}
