package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_updateAuditSuppressionCmd = &cobra.Command{
	Use:   "update-audit-suppression",
	Short: "Updates a Device Defender audit suppression.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_updateAuditSuppressionCmd).Standalone()

	iot_updateAuditSuppressionCmd.Flags().String("check-name", "", "")
	iot_updateAuditSuppressionCmd.Flags().String("description", "", "The description of the audit suppression.")
	iot_updateAuditSuppressionCmd.Flags().String("expiration-date", "", "The expiration date (epoch timestamp in seconds) that you want the suppression to adhere to.")
	iot_updateAuditSuppressionCmd.Flags().String("resource-identifier", "", "")
	iot_updateAuditSuppressionCmd.Flags().String("suppress-indefinitely", "", "Indicates whether a suppression should exist indefinitely or not.")
	iot_updateAuditSuppressionCmd.MarkFlagRequired("check-name")
	iot_updateAuditSuppressionCmd.MarkFlagRequired("resource-identifier")
	iotCmd.AddCommand(iot_updateAuditSuppressionCmd)
}
