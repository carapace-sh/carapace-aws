package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_createOutpostResolverCmd = &cobra.Command{
	Use:   "create-outpost-resolver",
	Short: "Creates a Route\u00a053 Resolver on an Outpost.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_createOutpostResolverCmd).Standalone()

	route53resolver_createOutpostResolverCmd.Flags().String("creator-request-id", "", "A unique string that identifies the request and that allows failed requests to be retried without the risk of running the operation twice.")
	route53resolver_createOutpostResolverCmd.Flags().String("instance-count", "", "Number of Amazon EC2 instances for the Resolver on Outpost.")
	route53resolver_createOutpostResolverCmd.Flags().String("name", "", "A friendly name that lets you easily find a configuration in the Resolver dashboard in the Route\u00a053 console.")
	route53resolver_createOutpostResolverCmd.Flags().String("outpost-arn", "", "The Amazon Resource Name (ARN) of the Outpost.")
	route53resolver_createOutpostResolverCmd.Flags().String("preferred-instance-type", "", "The Amazon EC2 instance type.")
	route53resolver_createOutpostResolverCmd.Flags().String("tags", "", "A string that helps identify the Route\u00a053 Resolvers on Outpost.")
	route53resolver_createOutpostResolverCmd.MarkFlagRequired("creator-request-id")
	route53resolver_createOutpostResolverCmd.MarkFlagRequired("name")
	route53resolver_createOutpostResolverCmd.MarkFlagRequired("outpost-arn")
	route53resolver_createOutpostResolverCmd.MarkFlagRequired("preferred-instance-type")
	route53resolverCmd.AddCommand(route53resolver_createOutpostResolverCmd)
}
