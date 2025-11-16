package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_createHostedZoneCmd = &cobra.Command{
	Use:   "create-hosted-zone",
	Short: "Creates a new public or private hosted zone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_createHostedZoneCmd).Standalone()

	route53_createHostedZoneCmd.Flags().String("caller-reference", "", "A unique string that identifies the request and that allows failed `CreateHostedZone` requests to be retried without the risk of executing the operation twice.")
	route53_createHostedZoneCmd.Flags().String("delegation-set-id", "", "If you want to associate a reusable delegation set with this hosted zone, the ID that Amazon Route\u00a053 assigned to the reusable delegation set when you created it.")
	route53_createHostedZoneCmd.Flags().String("hosted-zone-config", "", "(Optional) A complex type that contains the following optional values:")
	route53_createHostedZoneCmd.Flags().String("name", "", "The name of the domain.")
	route53_createHostedZoneCmd.Flags().String("vpc", "", "(Private hosted zones only) A complex type that contains information about the Amazon VPC that you're associating with this hosted zone.")
	route53_createHostedZoneCmd.MarkFlagRequired("caller-reference")
	route53_createHostedZoneCmd.MarkFlagRequired("name")
	route53Cmd.AddCommand(route53_createHostedZoneCmd)
}
