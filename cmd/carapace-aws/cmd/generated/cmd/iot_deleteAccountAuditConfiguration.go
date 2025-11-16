package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deleteAccountAuditConfigurationCmd = &cobra.Command{
	Use:   "delete-account-audit-configuration",
	Short: "Restores the default settings for Device Defender audits for this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deleteAccountAuditConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_deleteAccountAuditConfigurationCmd).Standalone()

		iot_deleteAccountAuditConfigurationCmd.Flags().String("delete-scheduled-audits", "", "If true, all scheduled audits are deleted.")
	})
	iotCmd.AddCommand(iot_deleteAccountAuditConfigurationCmd)
}
