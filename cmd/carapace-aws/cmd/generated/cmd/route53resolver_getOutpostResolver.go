package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_getOutpostResolverCmd = &cobra.Command{
	Use:   "get-outpost-resolver",
	Short: "Gets information about a specified Resolver on the Outpost, such as its instance count and type, name, and the current status of the Resolver.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_getOutpostResolverCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_getOutpostResolverCmd).Standalone()

		route53resolver_getOutpostResolverCmd.Flags().String("id", "", "The ID of the Resolver on the Outpost.")
		route53resolver_getOutpostResolverCmd.MarkFlagRequired("id")
	})
	route53resolverCmd.AddCommand(route53resolver_getOutpostResolverCmd)
}
