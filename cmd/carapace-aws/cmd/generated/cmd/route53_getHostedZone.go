package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_getHostedZoneCmd = &cobra.Command{
	Use:   "get-hosted-zone",
	Short: "Gets information about a specified hosted zone including the four name servers assigned to the hosted zone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_getHostedZoneCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53_getHostedZoneCmd).Standalone()

		route53_getHostedZoneCmd.Flags().String("id", "", "The ID of the hosted zone that you want to get information about.")
		route53_getHostedZoneCmd.MarkFlagRequired("id")
	})
	route53Cmd.AddCommand(route53_getHostedZoneCmd)
}
