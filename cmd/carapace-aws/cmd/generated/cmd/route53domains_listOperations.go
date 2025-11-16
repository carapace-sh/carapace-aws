package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domains_listOperationsCmd = &cobra.Command{
	Use:   "list-operations",
	Short: "Returns information about all of the operations that return an operation ID and that have ever been performed on domains that were registered by the current account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domains_listOperationsCmd).Standalone()

	route53domains_listOperationsCmd.Flags().String("marker", "", "For an initial request for a list of operations, omit this element.")
	route53domains_listOperationsCmd.Flags().String("max-items", "", "Number of domains to be returned.")
	route53domains_listOperationsCmd.Flags().String("sort-by", "", "The sort type for returned values.")
	route53domains_listOperationsCmd.Flags().String("sort-order", "", "The sort order for returned values, either ascending or descending.")
	route53domains_listOperationsCmd.Flags().String("status", "", "The status of the operations.")
	route53domains_listOperationsCmd.Flags().String("submitted-since", "", "An optional parameter that lets you get information about all the operations that you submitted after a specified date and time.")
	route53domains_listOperationsCmd.Flags().String("type", "", "An arrays of the domains operation types.")
	route53domainsCmd.AddCommand(route53domains_listOperationsCmd)
}
