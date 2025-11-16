package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_updateDataProtectionSettingsCmd = &cobra.Command{
	Use:   "update-data-protection-settings",
	Short: "Updates data protection settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_updateDataProtectionSettingsCmd).Standalone()

	workspacesWeb_updateDataProtectionSettingsCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	workspacesWeb_updateDataProtectionSettingsCmd.Flags().String("data-protection-settings-arn", "", "The ARN of the data protection settings.")
	workspacesWeb_updateDataProtectionSettingsCmd.Flags().String("description", "", "The description of the data protection settings.")
	workspacesWeb_updateDataProtectionSettingsCmd.Flags().String("display-name", "", "The display name of the data protection settings.")
	workspacesWeb_updateDataProtectionSettingsCmd.Flags().String("inline-redaction-configuration", "", "The inline redaction configuration of the data protection settings that will be applied to all sessions.")
	workspacesWeb_updateDataProtectionSettingsCmd.MarkFlagRequired("data-protection-settings-arn")
	workspacesWebCmd.AddCommand(workspacesWeb_updateDataProtectionSettingsCmd)
}
