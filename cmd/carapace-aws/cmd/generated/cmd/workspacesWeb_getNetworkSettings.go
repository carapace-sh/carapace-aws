package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_getNetworkSettingsCmd = &cobra.Command{
	Use:   "get-network-settings",
	Short: "Gets the network settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_getNetworkSettingsCmd).Standalone()

	workspacesWeb_getNetworkSettingsCmd.Flags().String("network-settings-arn", "", "The ARN of the network settings.")
	workspacesWeb_getNetworkSettingsCmd.MarkFlagRequired("network-settings-arn")
	workspacesWebCmd.AddCommand(workspacesWeb_getNetworkSettingsCmd)
}
