package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_listUserAccessLoggingSettingsCmd = &cobra.Command{
	Use:   "list-user-access-logging-settings",
	Short: "Retrieves a list of user access logging settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_listUserAccessLoggingSettingsCmd).Standalone()

	workspacesWeb_listUserAccessLoggingSettingsCmd.Flags().String("max-results", "", "The maximum number of results to be included in the next page.")
	workspacesWeb_listUserAccessLoggingSettingsCmd.Flags().String("next-token", "", "The pagination token used to retrieve the next page of results for this operation.")
	workspacesWebCmd.AddCommand(workspacesWeb_listUserAccessLoggingSettingsCmd)
}
