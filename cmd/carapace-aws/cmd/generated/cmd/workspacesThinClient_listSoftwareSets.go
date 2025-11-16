package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesThinClient_listSoftwareSetsCmd = &cobra.Command{
	Use:   "list-software-sets",
	Short: "Returns a list of software sets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesThinClient_listSoftwareSetsCmd).Standalone()

	workspacesThinClient_listSoftwareSetsCmd.Flags().String("max-results", "", "The maximum number of results that are returned per call.")
	workspacesThinClient_listSoftwareSetsCmd.Flags().String("next-token", "", "If `nextToken` is returned, there are more results available.")
	workspacesThinClientCmd.AddCommand(workspacesThinClient_listSoftwareSetsCmd)
}
