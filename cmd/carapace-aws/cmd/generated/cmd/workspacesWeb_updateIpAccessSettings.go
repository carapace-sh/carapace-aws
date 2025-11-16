package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_updateIpAccessSettingsCmd = &cobra.Command{
	Use:   "update-ip-access-settings",
	Short: "Updates IP access settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_updateIpAccessSettingsCmd).Standalone()

	workspacesWeb_updateIpAccessSettingsCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	workspacesWeb_updateIpAccessSettingsCmd.Flags().String("description", "", "The description of the IP access settings.")
	workspacesWeb_updateIpAccessSettingsCmd.Flags().String("display-name", "", "The display name of the IP access settings.")
	workspacesWeb_updateIpAccessSettingsCmd.Flags().String("ip-access-settings-arn", "", "The ARN of the IP access settings.")
	workspacesWeb_updateIpAccessSettingsCmd.Flags().String("ip-rules", "", "The updated IP rules of the IP access settings.")
	workspacesWeb_updateIpAccessSettingsCmd.MarkFlagRequired("ip-access-settings-arn")
	workspacesWebCmd.AddCommand(workspacesWeb_updateIpAccessSettingsCmd)
}
