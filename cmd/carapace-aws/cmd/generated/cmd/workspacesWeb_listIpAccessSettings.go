package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_listIpAccessSettingsCmd = &cobra.Command{
	Use:   "list-ip-access-settings",
	Short: "Retrieves a list of IP access settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_listIpAccessSettingsCmd).Standalone()

	workspacesWeb_listIpAccessSettingsCmd.Flags().String("max-results", "", "The maximum number of results to be included in the next page.")
	workspacesWeb_listIpAccessSettingsCmd.Flags().String("next-token", "", "The pagination token used to retrieve the next page of results for this operation.")
	workspacesWebCmd.AddCommand(workspacesWeb_listIpAccessSettingsCmd)
}
