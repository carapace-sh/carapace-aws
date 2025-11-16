package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_createResolverEndpointCmd = &cobra.Command{
	Use:   "create-resolver-endpoint",
	Short: "Creates a Resolver endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_createResolverEndpointCmd).Standalone()

	route53resolver_createResolverEndpointCmd.Flags().String("creator-request-id", "", "A unique string that identifies the request and that allows failed requests to be retried without the risk of running the operation twice.")
	route53resolver_createResolverEndpointCmd.Flags().String("direction", "", "Specify the applicable value:")
	route53resolver_createResolverEndpointCmd.Flags().String("ip-addresses", "", "The subnets and IP addresses in your VPC that DNS queries originate from (for outbound endpoints) or that you forward DNS queries to (for inbound endpoints).")
	route53resolver_createResolverEndpointCmd.Flags().String("name", "", "A friendly name that lets you easily find a configuration in the Resolver dashboard in the Route 53 console.")
	route53resolver_createResolverEndpointCmd.Flags().String("outpost-arn", "", "The Amazon Resource Name (ARN) of the Outpost.")
	route53resolver_createResolverEndpointCmd.Flags().String("preferred-instance-type", "", "The instance type.")
	route53resolver_createResolverEndpointCmd.Flags().String("protocols", "", "The protocols you want to use for the endpoint.")
	route53resolver_createResolverEndpointCmd.Flags().String("resolver-endpoint-type", "", "For the endpoint type you can choose either IPv4, IPv6, or dual-stack.")
	route53resolver_createResolverEndpointCmd.Flags().String("security-group-ids", "", "The ID of one or more security groups that you want to use to control access to this VPC.")
	route53resolver_createResolverEndpointCmd.Flags().String("tags", "", "A list of the tag keys and values that you want to associate with the endpoint.")
	route53resolver_createResolverEndpointCmd.MarkFlagRequired("creator-request-id")
	route53resolver_createResolverEndpointCmd.MarkFlagRequired("direction")
	route53resolver_createResolverEndpointCmd.MarkFlagRequired("ip-addresses")
	route53resolver_createResolverEndpointCmd.MarkFlagRequired("security-group-ids")
	route53resolverCmd.AddCommand(route53resolver_createResolverEndpointCmd)
}
