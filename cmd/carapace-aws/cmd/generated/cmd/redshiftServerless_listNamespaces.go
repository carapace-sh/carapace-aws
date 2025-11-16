package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_listNamespacesCmd = &cobra.Command{
	Use:   "list-namespaces",
	Short: "Returns information about a list of specified namespaces.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_listNamespacesCmd).Standalone()

	redshiftServerless_listNamespacesCmd.Flags().String("max-results", "", "An optional parameter that specifies the maximum number of results to return.")
	redshiftServerless_listNamespacesCmd.Flags().String("next-token", "", "If your initial `ListNamespaces` operation returns a `nextToken`, you can include the returned `nextToken` in following `ListNamespaces` operations, which returns results in the next page.")
	redshiftServerlessCmd.AddCommand(redshiftServerless_listNamespacesCmd)
}
