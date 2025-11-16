package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_associateNetworkSettingsCmd = &cobra.Command{
	Use:   "associate-network-settings",
	Short: "Associates a network settings resource with a web portal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_associateNetworkSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_associateNetworkSettingsCmd).Standalone()

		workspacesWeb_associateNetworkSettingsCmd.Flags().String("network-settings-arn", "", "The ARN of the network settings.")
		workspacesWeb_associateNetworkSettingsCmd.Flags().String("portal-arn", "", "The ARN of the web portal.")
		workspacesWeb_associateNetworkSettingsCmd.MarkFlagRequired("network-settings-arn")
		workspacesWeb_associateNetworkSettingsCmd.MarkFlagRequired("portal-arn")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_associateNetworkSettingsCmd)
}
