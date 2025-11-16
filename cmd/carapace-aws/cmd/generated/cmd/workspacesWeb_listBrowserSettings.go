package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_listBrowserSettingsCmd = &cobra.Command{
	Use:   "list-browser-settings",
	Short: "Retrieves a list of browser settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_listBrowserSettingsCmd).Standalone()

	workspacesWeb_listBrowserSettingsCmd.Flags().String("max-results", "", "The maximum number of results to be included in the next page.")
	workspacesWeb_listBrowserSettingsCmd.Flags().String("next-token", "", "The pagination token used to retrieve the next page of results for this operation.")
	workspacesWebCmd.AddCommand(workspacesWeb_listBrowserSettingsCmd)
}
