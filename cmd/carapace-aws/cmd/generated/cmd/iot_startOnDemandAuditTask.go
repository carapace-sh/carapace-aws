package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_startOnDemandAuditTaskCmd = &cobra.Command{
	Use:   "start-on-demand-audit-task",
	Short: "Starts an on-demand Device Defender audit.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_startOnDemandAuditTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_startOnDemandAuditTaskCmd).Standalone()

		iot_startOnDemandAuditTaskCmd.Flags().String("target-check-names", "", "Which checks are performed during the audit.")
		iot_startOnDemandAuditTaskCmd.MarkFlagRequired("target-check-names")
	})
	iotCmd.AddCommand(iot_startOnDemandAuditTaskCmd)
}
