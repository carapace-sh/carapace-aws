package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceGroups_updateAccountSettingsCmd = &cobra.Command{
	Use:   "update-account-settings",
	Short: "Turns on or turns off optional features in Resource Groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceGroups_updateAccountSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourceGroups_updateAccountSettingsCmd).Standalone()

		resourceGroups_updateAccountSettingsCmd.Flags().String("group-lifecycle-events-desired-status", "", "Specifies whether you want to turn [group lifecycle events](https://docs.aws.amazon.com/ARG/latest/userguide/monitor-groups.html) on or off.")
	})
	resourceGroupsCmd.AddCommand(resourceGroups_updateAccountSettingsCmd)
}
