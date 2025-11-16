package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_createDataProtectionSettingsCmd = &cobra.Command{
	Use:   "create-data-protection-settings",
	Short: "Creates a data protection settings resource that can be associated with a web portal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_createDataProtectionSettingsCmd).Standalone()

	workspacesWeb_createDataProtectionSettingsCmd.Flags().String("additional-encryption-context", "", "Additional encryption context of the data protection settings.")
	workspacesWeb_createDataProtectionSettingsCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	workspacesWeb_createDataProtectionSettingsCmd.Flags().String("customer-managed-key", "", "The custom managed key of the data protection settings.")
	workspacesWeb_createDataProtectionSettingsCmd.Flags().String("description", "", "The description of the data protection settings.")
	workspacesWeb_createDataProtectionSettingsCmd.Flags().String("display-name", "", "The display name of the data protection settings.")
	workspacesWeb_createDataProtectionSettingsCmd.Flags().String("inline-redaction-configuration", "", "The inline redaction configuration of the data protection settings that will be applied to all sessions.")
	workspacesWeb_createDataProtectionSettingsCmd.Flags().String("tags", "", "The tags to add to the data protection settings resource.")
	workspacesWebCmd.AddCommand(workspacesWeb_createDataProtectionSettingsCmd)
}
