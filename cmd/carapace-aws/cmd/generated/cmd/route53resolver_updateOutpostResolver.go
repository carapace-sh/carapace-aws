package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_updateOutpostResolverCmd = &cobra.Command{
	Use:   "update-outpost-resolver",
	Short: "You can use `UpdateOutpostResolver` to update the instance count, type, or name of a Resolver on an Outpost.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_updateOutpostResolverCmd).Standalone()

	route53resolver_updateOutpostResolverCmd.Flags().String("id", "", "A unique string that identifies Resolver on an Outpost.")
	route53resolver_updateOutpostResolverCmd.Flags().String("instance-count", "", "The Amazon EC2 instance count for a Resolver on the Outpost.")
	route53resolver_updateOutpostResolverCmd.Flags().String("name", "", "Name of the Resolver on the Outpost.")
	route53resolver_updateOutpostResolverCmd.Flags().String("preferred-instance-type", "", "Amazon EC2 instance type.")
	route53resolver_updateOutpostResolverCmd.MarkFlagRequired("id")
	route53resolverCmd.AddCommand(route53resolver_updateOutpostResolverCmd)
}
