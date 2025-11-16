package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceGroups_getAccountSettingsCmd = &cobra.Command{
	Use:   "get-account-settings",
	Short: "Retrieves the current status of optional features in Resource Groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceGroups_getAccountSettingsCmd).Standalone()

	resourceGroupsCmd.AddCommand(resourceGroups_getAccountSettingsCmd)
}
