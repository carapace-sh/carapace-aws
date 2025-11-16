package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_listSessionLoggersCmd = &cobra.Command{
	Use:   "list-session-loggers",
	Short: "Lists all available session logger resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_listSessionLoggersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_listSessionLoggersCmd).Standalone()

		workspacesWeb_listSessionLoggersCmd.Flags().String("max-results", "", "The maximum number of results to be included in the next page.")
		workspacesWeb_listSessionLoggersCmd.Flags().String("next-token", "", "The pagination token used to retrieve the next page of results for this operation.")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_listSessionLoggersCmd)
}
