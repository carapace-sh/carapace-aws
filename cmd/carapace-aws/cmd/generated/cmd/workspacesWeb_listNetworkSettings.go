package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_listNetworkSettingsCmd = &cobra.Command{
	Use:   "list-network-settings",
	Short: "Retrieves a list of network settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_listNetworkSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_listNetworkSettingsCmd).Standalone()

		workspacesWeb_listNetworkSettingsCmd.Flags().String("max-results", "", "The maximum number of results to be included in the next page.")
		workspacesWeb_listNetworkSettingsCmd.Flags().String("next-token", "", "The pagination token used to retrieve the next page of results for this operation.")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_listNetworkSettingsCmd)
}
