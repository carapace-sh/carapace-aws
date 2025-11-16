package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_getDataProtectionSettingsCmd = &cobra.Command{
	Use:   "get-data-protection-settings",
	Short: "Gets the data protection settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_getDataProtectionSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_getDataProtectionSettingsCmd).Standalone()

		workspacesWeb_getDataProtectionSettingsCmd.Flags().String("data-protection-settings-arn", "", "The ARN of the data protection settings.")
		workspacesWeb_getDataProtectionSettingsCmd.MarkFlagRequired("data-protection-settings-arn")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_getDataProtectionSettingsCmd)
}
