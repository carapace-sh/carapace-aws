package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_getHostedZoneLimitCmd = &cobra.Command{
	Use:   "get-hosted-zone-limit",
	Short: "Gets the specified limit for a specified hosted zone, for example, the maximum number of records that you can create in the hosted zone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_getHostedZoneLimitCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53_getHostedZoneLimitCmd).Standalone()

		route53_getHostedZoneLimitCmd.Flags().String("hosted-zone-id", "", "The ID of the hosted zone that you want to get a limit for.")
		route53_getHostedZoneLimitCmd.Flags().String("type", "", "The limit that you want to get.")
		route53_getHostedZoneLimitCmd.MarkFlagRequired("hosted-zone-id")
		route53_getHostedZoneLimitCmd.MarkFlagRequired("type")
	})
	route53Cmd.AddCommand(route53_getHostedZoneLimitCmd)
}
