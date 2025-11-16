package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceExplorer2_listResourcesCmd = &cobra.Command{
	Use:   "list-resources",
	Short: "Returns a list of resources and their details that match the specified criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceExplorer2_listResourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourceExplorer2_listResourcesCmd).Standalone()

		resourceExplorer2_listResourcesCmd.Flags().String("filters", "", "An array of strings that specify which resources are included in the results of queries made using this view.")
		resourceExplorer2_listResourcesCmd.Flags().String("max-results", "", "The maximum number of results that you want included on each page of the response.")
		resourceExplorer2_listResourcesCmd.Flags().String("next-token", "", "The parameter for receiving additional results if you receive a `NextToken` response in a previous request.")
		resourceExplorer2_listResourcesCmd.Flags().String("view-arn", "", "Specifies the Amazon resource name (ARN) of the view to use for the query.")
	})
	resourceExplorer2Cmd.AddCommand(resourceExplorer2_listResourcesCmd)
}
