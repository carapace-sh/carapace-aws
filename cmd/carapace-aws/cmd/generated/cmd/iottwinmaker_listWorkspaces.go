package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_listWorkspacesCmd = &cobra.Command{
	Use:   "list-workspaces",
	Short: "Retrieves information about workspaces in the current account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_listWorkspacesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iottwinmaker_listWorkspacesCmd).Standalone()

		iottwinmaker_listWorkspacesCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
		iottwinmaker_listWorkspacesCmd.Flags().String("next-token", "", "The string that specifies the next page of results.")
	})
	iottwinmakerCmd.AddCommand(iottwinmaker_listWorkspacesCmd)
}
