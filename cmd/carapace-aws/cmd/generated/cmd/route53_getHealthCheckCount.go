package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_getHealthCheckCountCmd = &cobra.Command{
	Use:   "get-health-check-count",
	Short: "Retrieves the number of health checks that are associated with the current Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_getHealthCheckCountCmd).Standalone()

	route53Cmd.AddCommand(route53_getHealthCheckCountCmd)
}
