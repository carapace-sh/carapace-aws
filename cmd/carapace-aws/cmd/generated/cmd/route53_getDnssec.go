package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_getDnssecCmd = &cobra.Command{
	Use:   "get-dnssec",
	Short: "Returns information about DNSSEC for a specific hosted zone, including the key-signing keys (KSKs) in the hosted zone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_getDnssecCmd).Standalone()

	route53_getDnssecCmd.Flags().String("hosted-zone-id", "", "A unique string used to identify a hosted zone.")
	route53_getDnssecCmd.MarkFlagRequired("hosted-zone-id")
	route53Cmd.AddCommand(route53_getDnssecCmd)
}
