package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_disassociateNetworkSettingsCmd = &cobra.Command{
	Use:   "disassociate-network-settings",
	Short: "Disassociates network settings from a web portal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_disassociateNetworkSettingsCmd).Standalone()

	workspacesWeb_disassociateNetworkSettingsCmd.Flags().String("portal-arn", "", "The ARN of the web portal.")
	workspacesWeb_disassociateNetworkSettingsCmd.MarkFlagRequired("portal-arn")
	workspacesWebCmd.AddCommand(workspacesWeb_disassociateNetworkSettingsCmd)
}
