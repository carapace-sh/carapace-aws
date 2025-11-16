package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appflow_listFlowsCmd = &cobra.Command{
	Use:   "list-flows",
	Short: "Lists all of the flows associated with your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appflow_listFlowsCmd).Standalone()

	appflow_listFlowsCmd.Flags().String("max-results", "", "Specifies the maximum number of items that should be returned in the result set.")
	appflow_listFlowsCmd.Flags().String("next-token", "", "The pagination token for next page of data.")
	appflowCmd.AddCommand(appflow_listFlowsCmd)
}
