package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_getIpAccessSettingsCmd = &cobra.Command{
	Use:   "get-ip-access-settings",
	Short: "Gets the IP access settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_getIpAccessSettingsCmd).Standalone()

	workspacesWeb_getIpAccessSettingsCmd.Flags().String("ip-access-settings-arn", "", "The ARN of the IP access settings.")
	workspacesWeb_getIpAccessSettingsCmd.MarkFlagRequired("ip-access-settings-arn")
	workspacesWebCmd.AddCommand(workspacesWeb_getIpAccessSettingsCmd)
}
