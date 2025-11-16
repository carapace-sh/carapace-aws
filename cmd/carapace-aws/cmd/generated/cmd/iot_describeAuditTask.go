package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeAuditTaskCmd = &cobra.Command{
	Use:   "describe-audit-task",
	Short: "Gets information about a Device Defender audit.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeAuditTaskCmd).Standalone()

	iot_describeAuditTaskCmd.Flags().String("task-id", "", "The ID of the audit whose information you want to get.")
	iot_describeAuditTaskCmd.MarkFlagRequired("task-id")
	iotCmd.AddCommand(iot_describeAuditTaskCmd)
}
