package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_getHealthCheckLastFailureReasonCmd = &cobra.Command{
	Use:   "get-health-check-last-failure-reason",
	Short: "Gets the reason that a specified health check failed most recently.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_getHealthCheckLastFailureReasonCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53_getHealthCheckLastFailureReasonCmd).Standalone()

		route53_getHealthCheckLastFailureReasonCmd.Flags().String("health-check-id", "", "The ID for the health check for which you want the last failure reason.")
		route53_getHealthCheckLastFailureReasonCmd.MarkFlagRequired("health-check-id")
	})
	route53Cmd.AddCommand(route53_getHealthCheckLastFailureReasonCmd)
}
