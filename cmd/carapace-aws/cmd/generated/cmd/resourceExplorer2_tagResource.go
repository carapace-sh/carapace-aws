package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceExplorer2_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds one or more tag key and value pairs to an Amazon Web Services Resource Explorer view or index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceExplorer2_tagResourceCmd).Standalone()

	resourceExplorer2_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the view or index that you want to attach tags to.")
	resourceExplorer2_tagResourceCmd.Flags().String("tags", "", "A list of tag key and value pairs that you want to attach to the specified view or index.")
	resourceExplorer2_tagResourceCmd.MarkFlagRequired("resource-arn")
	resourceExplorer2Cmd.AddCommand(resourceExplorer2_tagResourceCmd)
}
