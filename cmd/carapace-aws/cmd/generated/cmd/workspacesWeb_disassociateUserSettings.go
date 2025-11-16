package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_disassociateUserSettingsCmd = &cobra.Command{
	Use:   "disassociate-user-settings",
	Short: "Disassociates user settings from a web portal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_disassociateUserSettingsCmd).Standalone()

	workspacesWeb_disassociateUserSettingsCmd.Flags().String("portal-arn", "", "The ARN of the web portal.")
	workspacesWeb_disassociateUserSettingsCmd.MarkFlagRequired("portal-arn")
	workspacesWebCmd.AddCommand(workspacesWeb_disassociateUserSettingsCmd)
}
