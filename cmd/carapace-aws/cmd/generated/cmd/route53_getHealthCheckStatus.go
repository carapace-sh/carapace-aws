package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_getHealthCheckStatusCmd = &cobra.Command{
	Use:   "get-health-check-status",
	Short: "Gets status of a specified health check.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_getHealthCheckStatusCmd).Standalone()

	route53_getHealthCheckStatusCmd.Flags().String("health-check-id", "", "The ID for the health check that you want the current status for.")
	route53_getHealthCheckStatusCmd.MarkFlagRequired("health-check-id")
	route53Cmd.AddCommand(route53_getHealthCheckStatusCmd)
}
