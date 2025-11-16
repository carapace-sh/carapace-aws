package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeAuditSuppressionCmd = &cobra.Command{
	Use:   "describe-audit-suppression",
	Short: "Gets information about a Device Defender audit suppression.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeAuditSuppressionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_describeAuditSuppressionCmd).Standalone()

		iot_describeAuditSuppressionCmd.Flags().String("check-name", "", "")
		iot_describeAuditSuppressionCmd.Flags().String("resource-identifier", "", "")
		iot_describeAuditSuppressionCmd.MarkFlagRequired("check-name")
		iot_describeAuditSuppressionCmd.MarkFlagRequired("resource-identifier")
	})
	iotCmd.AddCommand(iot_describeAuditSuppressionCmd)
}
