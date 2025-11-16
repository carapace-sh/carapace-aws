package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_getHostedZoneCountCmd = &cobra.Command{
	Use:   "get-hosted-zone-count",
	Short: "Retrieves the number of hosted zones that are associated with the current Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_getHostedZoneCountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53_getHostedZoneCountCmd).Standalone()

	})
	route53Cmd.AddCommand(route53_getHostedZoneCountCmd)
}
