package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_listOutpostResolversCmd = &cobra.Command{
	Use:   "list-outpost-resolvers",
	Short: "Lists all the Resolvers on Outposts that were created using the current Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_listOutpostResolversCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_listOutpostResolversCmd).Standalone()

		route53resolver_listOutpostResolversCmd.Flags().String("max-results", "", "The maximum number of Resolvers on the Outpost that you want to return in the response to a `ListOutpostResolver` request.")
		route53resolver_listOutpostResolversCmd.Flags().String("next-token", "", "For the first `ListOutpostResolver` request, omit this value.")
		route53resolver_listOutpostResolversCmd.Flags().String("outpost-arn", "", "The Amazon Resource Name (ARN) of the Outpost.")
	})
	route53resolverCmd.AddCommand(route53resolver_listOutpostResolversCmd)
}
