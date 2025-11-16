package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_deleteOutpostResolverCmd = &cobra.Command{
	Use:   "delete-outpost-resolver",
	Short: "Deletes a Resolver on the Outpost.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_deleteOutpostResolverCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_deleteOutpostResolverCmd).Standalone()

		route53resolver_deleteOutpostResolverCmd.Flags().String("id", "", "A unique string that identifies the Resolver on the Outpost.")
		route53resolver_deleteOutpostResolverCmd.MarkFlagRequired("id")
	})
	route53resolverCmd.AddCommand(route53resolver_deleteOutpostResolverCmd)
}
