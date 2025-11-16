package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_disassociateIpAccessSettingsCmd = &cobra.Command{
	Use:   "disassociate-ip-access-settings",
	Short: "Disassociates IP access settings from a web portal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_disassociateIpAccessSettingsCmd).Standalone()

	workspacesWeb_disassociateIpAccessSettingsCmd.Flags().String("portal-arn", "", "The ARN of the web portal.")
	workspacesWeb_disassociateIpAccessSettingsCmd.MarkFlagRequired("portal-arn")
	workspacesWebCmd.AddCommand(workspacesWeb_disassociateIpAccessSettingsCmd)
}
