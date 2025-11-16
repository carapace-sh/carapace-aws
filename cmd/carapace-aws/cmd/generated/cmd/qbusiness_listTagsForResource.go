package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Gets a list of tags associated with a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_listTagsForResourceCmd).Standalone()

		qbusiness_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Amazon Q Business application or data source to get a list of tags for.")
		qbusiness_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	qbusinessCmd.AddCommand(qbusiness_listTagsForResourceCmd)
}
