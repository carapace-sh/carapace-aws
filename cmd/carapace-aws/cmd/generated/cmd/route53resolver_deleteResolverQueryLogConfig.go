package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_deleteResolverQueryLogConfigCmd = &cobra.Command{
	Use:   "delete-resolver-query-log-config",
	Short: "Deletes a query logging configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_deleteResolverQueryLogConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53resolver_deleteResolverQueryLogConfigCmd).Standalone()

		route53resolver_deleteResolverQueryLogConfigCmd.Flags().String("resolver-query-log-config-id", "", "The ID of the query logging configuration that you want to delete.")
		route53resolver_deleteResolverQueryLogConfigCmd.MarkFlagRequired("resolver-query-log-config-id")
	})
	route53resolverCmd.AddCommand(route53resolver_deleteResolverQueryLogConfigCmd)
}
