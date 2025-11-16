package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_executeQueryCmd = &cobra.Command{
	Use:   "execute-query",
	Short: "Run queries to access information from your knowledge graph of entities within individual workspaces.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_executeQueryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iottwinmaker_executeQueryCmd).Standalone()

		iottwinmaker_executeQueryCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
		iottwinmaker_executeQueryCmd.Flags().String("next-token", "", "The string that specifies the next page of results.")
		iottwinmaker_executeQueryCmd.Flags().String("query-statement", "", "The query statement.")
		iottwinmaker_executeQueryCmd.Flags().String("workspace-id", "", "The ID of the workspace.")
		iottwinmaker_executeQueryCmd.MarkFlagRequired("query-statement")
		iottwinmaker_executeQueryCmd.MarkFlagRequired("workspace-id")
	})
	iottwinmakerCmd.AddCommand(iottwinmaker_executeQueryCmd)
}
