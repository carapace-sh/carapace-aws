package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_listHostedZonesCmd = &cobra.Command{
	Use:   "list-hosted-zones",
	Short: "Retrieves a list of the public and private hosted zones that are associated with the current Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_listHostedZonesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53_listHostedZonesCmd).Standalone()

		route53_listHostedZonesCmd.Flags().String("delegation-set-id", "", "If you're using reusable delegation sets and you want to list all of the hosted zones that are associated with a reusable delegation set, specify the ID of that reusable delegation set.")
		route53_listHostedZonesCmd.Flags().String("hosted-zone-type", "", "(Optional) Specifies if the hosted zone is private.")
		route53_listHostedZonesCmd.Flags().String("marker", "", "If the value of `IsTruncated` in the previous response was `true`, you have more hosted zones.")
		route53_listHostedZonesCmd.Flags().String("max-items", "", "(Optional) The maximum number of hosted zones that you want Amazon Route 53 to return.")
	})
	route53Cmd.AddCommand(route53_listHostedZonesCmd)
}
