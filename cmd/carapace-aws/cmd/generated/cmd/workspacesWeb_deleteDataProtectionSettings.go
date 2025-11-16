package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_deleteDataProtectionSettingsCmd = &cobra.Command{
	Use:   "delete-data-protection-settings",
	Short: "Deletes data protection settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_deleteDataProtectionSettingsCmd).Standalone()

	workspacesWeb_deleteDataProtectionSettingsCmd.Flags().String("data-protection-settings-arn", "", "The ARN of the data protection settings.")
	workspacesWeb_deleteDataProtectionSettingsCmd.MarkFlagRequired("data-protection-settings-arn")
	workspacesWebCmd.AddCommand(workspacesWeb_deleteDataProtectionSettingsCmd)
}
