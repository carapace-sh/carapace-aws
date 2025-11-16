package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_listSessionsCmd = &cobra.Command{
	Use:   "list-sessions",
	Short: "Lists information for multiple secure browser sessions from a specific portal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_listSessionsCmd).Standalone()

	workspacesWeb_listSessionsCmd.Flags().String("max-results", "", "The maximum number of results to be included in the next page.")
	workspacesWeb_listSessionsCmd.Flags().String("next-token", "", "The pagination token used to retrieve the next page of results for this operation.")
	workspacesWeb_listSessionsCmd.Flags().String("portal-id", "", "The ID of the web portal for the sessions.")
	workspacesWeb_listSessionsCmd.Flags().String("session-id", "", "The ID of the session.")
	workspacesWeb_listSessionsCmd.Flags().String("sort-by", "", "The method in which the returned sessions should be sorted.")
	workspacesWeb_listSessionsCmd.Flags().String("status", "", "The status of the session.")
	workspacesWeb_listSessionsCmd.Flags().String("username", "", "The username of the session.")
	workspacesWeb_listSessionsCmd.MarkFlagRequired("portal-id")
	workspacesWebCmd.AddCommand(workspacesWeb_listSessionsCmd)
}
