package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_listHostedZonesByNameCmd = &cobra.Command{
	Use:   "list-hosted-zones-by-name",
	Short: "Retrieves a list of your hosted zones in lexicographic order.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_listHostedZonesByNameCmd).Standalone()

	route53_listHostedZonesByNameCmd.Flags().String("dnsname", "", "(Optional) For your first request to `ListHostedZonesByName`, include the `dnsname` parameter only if you want to specify the name of the first hosted zone in the response.")
	route53_listHostedZonesByNameCmd.Flags().String("hosted-zone-id", "", "(Optional) For your first request to `ListHostedZonesByName`, do not include the `hostedzoneid` parameter.")
	route53_listHostedZonesByNameCmd.Flags().String("max-items", "", "The maximum number of hosted zones to be included in the response body for this request.")
	route53Cmd.AddCommand(route53_listHostedZonesByNameCmd)
}
