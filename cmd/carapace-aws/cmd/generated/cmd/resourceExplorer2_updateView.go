package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceExplorer2_updateViewCmd = &cobra.Command{
	Use:   "update-view",
	Short: "Modifies some of the details of a view.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceExplorer2_updateViewCmd).Standalone()

	resourceExplorer2_updateViewCmd.Flags().String("filters", "", "An array of strings that specify which resources are included in the results of queries made using this view.")
	resourceExplorer2_updateViewCmd.Flags().String("included-properties", "", "Specifies optional fields that you want included in search results from this view.")
	resourceExplorer2_updateViewCmd.Flags().String("view-arn", "", "The [Amazon resource name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the view that you want to modify.")
	resourceExplorer2_updateViewCmd.MarkFlagRequired("view-arn")
	resourceExplorer2Cmd.AddCommand(resourceExplorer2_updateViewCmd)
}
