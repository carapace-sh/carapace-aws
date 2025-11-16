package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deleteAuditSuppressionCmd = &cobra.Command{
	Use:   "delete-audit-suppression",
	Short: "Deletes a Device Defender audit suppression.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deleteAuditSuppressionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_deleteAuditSuppressionCmd).Standalone()

		iot_deleteAuditSuppressionCmd.Flags().String("check-name", "", "")
		iot_deleteAuditSuppressionCmd.Flags().String("resource-identifier", "", "")
		iot_deleteAuditSuppressionCmd.MarkFlagRequired("check-name")
		iot_deleteAuditSuppressionCmd.MarkFlagRequired("resource-identifier")
	})
	iotCmd.AddCommand(iot_deleteAuditSuppressionCmd)
}
