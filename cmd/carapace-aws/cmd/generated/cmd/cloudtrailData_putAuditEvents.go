package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrailData_putAuditEventsCmd = &cobra.Command{
	Use:   "put-audit-events",
	Short: "Ingests your application events into CloudTrail Lake.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrailData_putAuditEventsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudtrailData_putAuditEventsCmd).Standalone()

		cloudtrailData_putAuditEventsCmd.Flags().String("audit-events", "", "The JSON payload of events that you want to ingest.")
		cloudtrailData_putAuditEventsCmd.Flags().String("channel-arn", "", "The ARN or ID (the ARN suffix) of a channel.")
		cloudtrailData_putAuditEventsCmd.Flags().String("external-id", "", "A unique identifier that is conditionally required when the channel's resource policy includes an external ID.")
		cloudtrailData_putAuditEventsCmd.MarkFlagRequired("audit-events")
		cloudtrailData_putAuditEventsCmd.MarkFlagRequired("channel-arn")
	})
	cloudtrailDataCmd.AddCommand(cloudtrailData_putAuditEventsCmd)
}
