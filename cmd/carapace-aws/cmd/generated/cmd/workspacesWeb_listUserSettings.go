package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_listUserSettingsCmd = &cobra.Command{
	Use:   "list-user-settings",
	Short: "Retrieves a list of user settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_listUserSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_listUserSettingsCmd).Standalone()

		workspacesWeb_listUserSettingsCmd.Flags().String("max-results", "", "The maximum number of results to be included in the next page.")
		workspacesWeb_listUserSettingsCmd.Flags().String("next-token", "", "The pagination token used to retrieve the next page of results for this operation.")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_listUserSettingsCmd)
}
