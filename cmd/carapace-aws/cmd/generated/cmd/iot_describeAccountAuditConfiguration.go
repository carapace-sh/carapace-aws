package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeAccountAuditConfigurationCmd = &cobra.Command{
	Use:   "describe-account-audit-configuration",
	Short: "Gets information about the Device Defender audit settings for this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeAccountAuditConfigurationCmd).Standalone()

	iotCmd.AddCommand(iot_describeAccountAuditConfigurationCmd)
}
