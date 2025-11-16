package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceExplorer2_listViewsCmd = &cobra.Command{
	Use:   "list-views",
	Short: "Lists the [Amazon resource names (ARNs)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the views available in the Amazon Web Services Region in which you call this operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceExplorer2_listViewsCmd).Standalone()

	resourceExplorer2_listViewsCmd.Flags().String("max-results", "", "The maximum number of results that you want included on each page of the response.")
	resourceExplorer2_listViewsCmd.Flags().String("next-token", "", "The parameter for receiving additional results if you receive a `NextToken` response in a previous request.")
	resourceExplorer2Cmd.AddCommand(resourceExplorer2_listViewsCmd)
}
