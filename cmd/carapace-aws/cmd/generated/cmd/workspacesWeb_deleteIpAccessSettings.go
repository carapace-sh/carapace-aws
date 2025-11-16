package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_deleteIpAccessSettingsCmd = &cobra.Command{
	Use:   "delete-ip-access-settings",
	Short: "Deletes IP access settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_deleteIpAccessSettingsCmd).Standalone()

	workspacesWeb_deleteIpAccessSettingsCmd.Flags().String("ip-access-settings-arn", "", "The ARN of the IP access settings.")
	workspacesWeb_deleteIpAccessSettingsCmd.MarkFlagRequired("ip-access-settings-arn")
	workspacesWebCmd.AddCommand(workspacesWeb_deleteIpAccessSettingsCmd)
}
