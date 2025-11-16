package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_associateIpAccessSettingsCmd = &cobra.Command{
	Use:   "associate-ip-access-settings",
	Short: "Associates an IP access settings resource with a web portal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_associateIpAccessSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_associateIpAccessSettingsCmd).Standalone()

		workspacesWeb_associateIpAccessSettingsCmd.Flags().String("ip-access-settings-arn", "", "The ARN of the IP access settings.")
		workspacesWeb_associateIpAccessSettingsCmd.Flags().String("portal-arn", "", "The ARN of the web portal.")
		workspacesWeb_associateIpAccessSettingsCmd.MarkFlagRequired("ip-access-settings-arn")
		workspacesWeb_associateIpAccessSettingsCmd.MarkFlagRequired("portal-arn")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_associateIpAccessSettingsCmd)
}
