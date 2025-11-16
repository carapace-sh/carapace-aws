package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesThinClient_listEnvironmentsCmd = &cobra.Command{
	Use:   "list-environments",
	Short: "Returns a list of environments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesThinClient_listEnvironmentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesThinClient_listEnvironmentsCmd).Standalone()

		workspacesThinClient_listEnvironmentsCmd.Flags().String("max-results", "", "The maximum number of results that are returned per call.")
		workspacesThinClient_listEnvironmentsCmd.Flags().String("next-token", "", "If `nextToken` is returned, there are more results available.")
	})
	workspacesThinClientCmd.AddCommand(workspacesThinClient_listEnvironmentsCmd)
}
