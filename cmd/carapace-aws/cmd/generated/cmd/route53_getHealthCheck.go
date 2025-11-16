package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_getHealthCheckCmd = &cobra.Command{
	Use:   "get-health-check",
	Short: "Gets information about a specified health check.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_getHealthCheckCmd).Standalone()

	route53_getHealthCheckCmd.Flags().String("health-check-id", "", "The identifier that Amazon Route 53 assigned to the health check when you created it.")
	route53_getHealthCheckCmd.MarkFlagRequired("health-check-id")
	route53Cmd.AddCommand(route53_getHealthCheckCmd)
}
