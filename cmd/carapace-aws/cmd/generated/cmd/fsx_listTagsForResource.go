package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists tags for Amazon FSx resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fsx_listTagsForResourceCmd).Standalone()

		fsx_listTagsForResourceCmd.Flags().String("max-results", "", "Maximum number of tags to return in the response (integer).")
		fsx_listTagsForResourceCmd.Flags().String("next-token", "", "Opaque pagination token returned from a previous `ListTagsForResource` operation (String).")
		fsx_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the Amazon FSx resource that will have its tags listed.")
		fsx_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	fsxCmd.AddCommand(fsx_listTagsForResourceCmd)
}
