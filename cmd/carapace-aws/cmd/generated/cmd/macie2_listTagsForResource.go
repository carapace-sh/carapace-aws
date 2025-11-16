package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Retrieves the tags (keys and values) that are associated with an Amazon Macie resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_listTagsForResourceCmd).Standalone()

		macie2_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		macie2_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	macie2Cmd.AddCommand(macie2_listTagsForResourceCmd)
}
