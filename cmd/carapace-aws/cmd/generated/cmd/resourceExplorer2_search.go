package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceExplorer2_searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Searches for resources and displays details about all resources that match the specified criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceExplorer2_searchCmd).Standalone()

	resourceExplorer2_searchCmd.Flags().String("max-results", "", "The maximum number of results that you want included on each page of the response.")
	resourceExplorer2_searchCmd.Flags().String("next-token", "", "The parameter for receiving additional results if you receive a `NextToken` response in a previous request.")
	resourceExplorer2_searchCmd.Flags().String("query-string", "", "A string that includes keywords and filters that specify the resources that you want to include in the results.")
	resourceExplorer2_searchCmd.Flags().String("view-arn", "", "Specifies the [Amazon resource name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the view to use for the query.")
	resourceExplorer2_searchCmd.MarkFlagRequired("query-string")
	resourceExplorer2Cmd.AddCommand(resourceExplorer2_searchCmd)
}
