package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeScheduledAuditCmd = &cobra.Command{
	Use:   "describe-scheduled-audit",
	Short: "Gets information about a scheduled audit.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeScheduledAuditCmd).Standalone()

	iot_describeScheduledAuditCmd.Flags().String("scheduled-audit-name", "", "The name of the scheduled audit whose information you want to get.")
	iot_describeScheduledAuditCmd.MarkFlagRequired("scheduled-audit-name")
	iotCmd.AddCommand(iot_describeScheduledAuditCmd)
}
