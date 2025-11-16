package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_createAuditSuppressionCmd = &cobra.Command{
	Use:   "create-audit-suppression",
	Short: "Creates a Device Defender audit suppression.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_createAuditSuppressionCmd).Standalone()

	iot_createAuditSuppressionCmd.Flags().String("check-name", "", "")
	iot_createAuditSuppressionCmd.Flags().String("client-request-token", "", "Each audit supression must have a unique client request token.")
	iot_createAuditSuppressionCmd.Flags().String("description", "", "The description of the audit suppression.")
	iot_createAuditSuppressionCmd.Flags().String("expiration-date", "", "The epoch timestamp in seconds at which this suppression expires.")
	iot_createAuditSuppressionCmd.Flags().String("resource-identifier", "", "")
	iot_createAuditSuppressionCmd.Flags().String("suppress-indefinitely", "", "Indicates whether a suppression should exist indefinitely or not.")
	iot_createAuditSuppressionCmd.MarkFlagRequired("check-name")
	iot_createAuditSuppressionCmd.MarkFlagRequired("client-request-token")
	iot_createAuditSuppressionCmd.MarkFlagRequired("resource-identifier")
	iotCmd.AddCommand(iot_createAuditSuppressionCmd)
}
