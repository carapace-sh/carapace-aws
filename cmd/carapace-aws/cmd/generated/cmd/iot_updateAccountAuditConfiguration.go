package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_updateAccountAuditConfigurationCmd = &cobra.Command{
	Use:   "update-account-audit-configuration",
	Short: "Configures or reconfigures the Device Defender audit settings for this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_updateAccountAuditConfigurationCmd).Standalone()

	iot_updateAccountAuditConfigurationCmd.Flags().String("audit-check-configurations", "", "Specifies which audit checks are enabled and disabled for this account.")
	iot_updateAccountAuditConfigurationCmd.Flags().String("audit-notification-target-configurations", "", "Information about the targets to which audit notifications are sent.")
	iot_updateAccountAuditConfigurationCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the role that grants permission to IoT to access information about your devices, policies, certificates, and other items as required when performing an audit.")
	iotCmd.AddCommand(iot_updateAccountAuditConfigurationCmd)
}
