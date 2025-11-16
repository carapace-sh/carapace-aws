package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicediscovery_listNamespacesCmd = &cobra.Command{
	Use:   "list-namespaces",
	Short: "Lists summary information about the namespaces that were created by the current Amazon Web Services account and shared with the current Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicediscovery_listNamespacesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicediscovery_listNamespacesCmd).Standalone()

		servicediscovery_listNamespacesCmd.Flags().String("filters", "", "A complex type that contains specifications for the namespaces that you want to list.")
		servicediscovery_listNamespacesCmd.Flags().String("max-results", "", "The maximum number of namespaces that you want Cloud Map to return in the response to a `ListNamespaces` request.")
		servicediscovery_listNamespacesCmd.Flags().String("next-token", "", "For the first `ListNamespaces` request, omit this value.")
	})
	servicediscoveryCmd.AddCommand(servicediscovery_listNamespacesCmd)
}
