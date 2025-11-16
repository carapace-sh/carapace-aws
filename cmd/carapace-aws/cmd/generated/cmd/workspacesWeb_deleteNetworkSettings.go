package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_deleteNetworkSettingsCmd = &cobra.Command{
	Use:   "delete-network-settings",
	Short: "Deletes network settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_deleteNetworkSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_deleteNetworkSettingsCmd).Standalone()

		workspacesWeb_deleteNetworkSettingsCmd.Flags().String("network-settings-arn", "", "The ARN of the network settings.")
		workspacesWeb_deleteNetworkSettingsCmd.MarkFlagRequired("network-settings-arn")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_deleteNetworkSettingsCmd)
}
