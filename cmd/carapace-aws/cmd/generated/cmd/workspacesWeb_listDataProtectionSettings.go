package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_listDataProtectionSettingsCmd = &cobra.Command{
	Use:   "list-data-protection-settings",
	Short: "Retrieves a list of data protection settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_listDataProtectionSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_listDataProtectionSettingsCmd).Standalone()

		workspacesWeb_listDataProtectionSettingsCmd.Flags().String("max-results", "", "The maximum number of results to be included in the next page.")
		workspacesWeb_listDataProtectionSettingsCmd.Flags().String("next-token", "", "The pagination token used to retrieve the next page of results for this operation.")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_listDataProtectionSettingsCmd)
}
