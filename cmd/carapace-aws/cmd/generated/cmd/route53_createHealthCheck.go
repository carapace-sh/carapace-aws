package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_createHealthCheckCmd = &cobra.Command{
	Use:   "create-health-check",
	Short: "Creates a new health check.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_createHealthCheckCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53_createHealthCheckCmd).Standalone()

		route53_createHealthCheckCmd.Flags().String("caller-reference", "", "A unique string that identifies the request and that allows you to retry a failed `CreateHealthCheck` request without the risk of creating two identical health checks:")
		route53_createHealthCheckCmd.Flags().String("health-check-config", "", "A complex type that contains settings for a new health check.")
		route53_createHealthCheckCmd.MarkFlagRequired("caller-reference")
		route53_createHealthCheckCmd.MarkFlagRequired("health-check-config")
	})
	route53Cmd.AddCommand(route53_createHealthCheckCmd)
}
