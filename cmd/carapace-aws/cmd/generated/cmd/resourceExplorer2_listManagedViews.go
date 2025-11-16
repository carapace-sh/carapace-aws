package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceExplorer2_listManagedViewsCmd = &cobra.Command{
	Use:   "list-managed-views",
	Short: "Lists the Amazon resource names (ARNs) of the [Amazon Web Services-managed views](https://docs.aws.amazon.com/resource-explorer/latest/userguide/aws-managed-views.html) available in the Amazon Web Services Region in which you call this operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceExplorer2_listManagedViewsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourceExplorer2_listManagedViewsCmd).Standalone()

		resourceExplorer2_listManagedViewsCmd.Flags().String("max-results", "", "The maximum number of results that you want included on each page of the response.")
		resourceExplorer2_listManagedViewsCmd.Flags().String("next-token", "", "The parameter for receiving additional results if you receive a `NextToken` response in a previous request.")
		resourceExplorer2_listManagedViewsCmd.Flags().String("service-principal", "", "Specifies a service principal name.")
	})
	resourceExplorer2Cmd.AddCommand(resourceExplorer2_listManagedViewsCmd)
}
