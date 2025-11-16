package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceExplorer2_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags that are attached to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceExplorer2_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourceExplorer2_listTagsForResourceCmd).Standalone()

		resourceExplorer2_listTagsForResourceCmd.Flags().String("resource-arn", "", "The [Amazon resource name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the view or index that you want to attach tags to.")
		resourceExplorer2_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	resourceExplorer2Cmd.AddCommand(resourceExplorer2_listTagsForResourceCmd)
}
