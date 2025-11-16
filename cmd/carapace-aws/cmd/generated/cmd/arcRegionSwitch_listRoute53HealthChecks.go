package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arcRegionSwitch_listRoute53HealthChecksCmd = &cobra.Command{
	Use:   "list-route53-health-checks",
	Short: "List the Amazon Route 53 health checks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(arcRegionSwitch_listRoute53HealthChecksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(arcRegionSwitch_listRoute53HealthChecksCmd).Standalone()

		arcRegionSwitch_listRoute53HealthChecksCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the Amazon Route 53 health check request.")
		arcRegionSwitch_listRoute53HealthChecksCmd.Flags().String("hosted-zone-id", "", "The hosted zone ID for the health checks.")
		arcRegionSwitch_listRoute53HealthChecksCmd.Flags().String("max-results", "", "The number of objects that you want to return with this call.")
		arcRegionSwitch_listRoute53HealthChecksCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
		arcRegionSwitch_listRoute53HealthChecksCmd.Flags().String("record-name", "", "The record name for the health checks.")
		arcRegionSwitch_listRoute53HealthChecksCmd.MarkFlagRequired("arn")
	})
	arcRegionSwitchCmd.AddCommand(arcRegionSwitch_listRoute53HealthChecksCmd)
}
