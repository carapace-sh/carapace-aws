package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_createIpAccessSettingsCmd = &cobra.Command{
	Use:   "create-ip-access-settings",
	Short: "Creates an IP access settings resource that can be associated with a web portal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_createIpAccessSettingsCmd).Standalone()

	workspacesWeb_createIpAccessSettingsCmd.Flags().String("additional-encryption-context", "", "Additional encryption context of the IP access settings.")
	workspacesWeb_createIpAccessSettingsCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	workspacesWeb_createIpAccessSettingsCmd.Flags().String("customer-managed-key", "", "The custom managed key of the IP access settings.")
	workspacesWeb_createIpAccessSettingsCmd.Flags().String("description", "", "The description of the IP access settings.")
	workspacesWeb_createIpAccessSettingsCmd.Flags().String("display-name", "", "The display name of the IP access settings.")
	workspacesWeb_createIpAccessSettingsCmd.Flags().String("ip-rules", "", "The IP rules of the IP access settings.")
	workspacesWeb_createIpAccessSettingsCmd.Flags().String("tags", "", "The tags to add to the IP access settings resource.")
	workspacesWeb_createIpAccessSettingsCmd.MarkFlagRequired("ip-rules")
	workspacesWebCmd.AddCommand(workspacesWeb_createIpAccessSettingsCmd)
}
