package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_associateDataProtectionSettingsCmd = &cobra.Command{
	Use:   "associate-data-protection-settings",
	Short: "Associates a data protection settings resource with a web portal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_associateDataProtectionSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_associateDataProtectionSettingsCmd).Standalone()

		workspacesWeb_associateDataProtectionSettingsCmd.Flags().String("data-protection-settings-arn", "", "The ARN of the data protection settings.")
		workspacesWeb_associateDataProtectionSettingsCmd.Flags().String("portal-arn", "", "The ARN of the web portal.")
		workspacesWeb_associateDataProtectionSettingsCmd.MarkFlagRequired("data-protection-settings-arn")
		workspacesWeb_associateDataProtectionSettingsCmd.MarkFlagRequired("portal-arn")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_associateDataProtectionSettingsCmd)
}
