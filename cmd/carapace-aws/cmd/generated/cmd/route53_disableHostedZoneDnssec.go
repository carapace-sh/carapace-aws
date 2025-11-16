package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_disableHostedZoneDnssecCmd = &cobra.Command{
	Use:   "disable-hosted-zone-dnssec",
	Short: "Disables DNSSEC signing in a specific hosted zone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_disableHostedZoneDnssecCmd).Standalone()

	route53_disableHostedZoneDnssecCmd.Flags().String("hosted-zone-id", "", "A unique string used to identify a hosted zone.")
	route53_disableHostedZoneDnssecCmd.MarkFlagRequired("hosted-zone-id")
	route53Cmd.AddCommand(route53_disableHostedZoneDnssecCmd)
}
