package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_disassociateDataProtectionSettingsCmd = &cobra.Command{
	Use:   "disassociate-data-protection-settings",
	Short: "Disassociates data protection settings from a web portal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_disassociateDataProtectionSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_disassociateDataProtectionSettingsCmd).Standalone()

		workspacesWeb_disassociateDataProtectionSettingsCmd.Flags().String("portal-arn", "", "The ARN of the web portal.")
		workspacesWeb_disassociateDataProtectionSettingsCmd.MarkFlagRequired("portal-arn")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_disassociateDataProtectionSettingsCmd)
}
