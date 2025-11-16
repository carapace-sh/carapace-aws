package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_disassociateUserAccessLoggingSettingsCmd = &cobra.Command{
	Use:   "disassociate-user-access-logging-settings",
	Short: "Disassociates user access logging settings from a web portal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_disassociateUserAccessLoggingSettingsCmd).Standalone()

	workspacesWeb_disassociateUserAccessLoggingSettingsCmd.Flags().String("portal-arn", "", "The ARN of the web portal.")
	workspacesWeb_disassociateUserAccessLoggingSettingsCmd.MarkFlagRequired("portal-arn")
	workspacesWebCmd.AddCommand(workspacesWeb_disassociateUserAccessLoggingSettingsCmd)
}
