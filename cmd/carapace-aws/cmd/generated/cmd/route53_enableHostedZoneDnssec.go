package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_enableHostedZoneDnssecCmd = &cobra.Command{
	Use:   "enable-hosted-zone-dnssec",
	Short: "Enables DNSSEC signing in a specific hosted zone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_enableHostedZoneDnssecCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53_enableHostedZoneDnssecCmd).Standalone()

		route53_enableHostedZoneDnssecCmd.Flags().String("hosted-zone-id", "", "A unique string used to identify a hosted zone.")
		route53_enableHostedZoneDnssecCmd.MarkFlagRequired("hosted-zone-id")
	})
	route53Cmd.AddCommand(route53_enableHostedZoneDnssecCmd)
}
