package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_updateQuerySuggestionsBlockListCmd = &cobra.Command{
	Use:   "update-query-suggestions-block-list",
	Short: "Updates a block list used for query suggestions for an index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_updateQuerySuggestionsBlockListCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendra_updateQuerySuggestionsBlockListCmd).Standalone()

		kendra_updateQuerySuggestionsBlockListCmd.Flags().String("description", "", "A new description for the block list.")
		kendra_updateQuerySuggestionsBlockListCmd.Flags().String("id", "", "The identifier of the block list you want to update.")
		kendra_updateQuerySuggestionsBlockListCmd.Flags().String("index-id", "", "The identifier of the index for the block list.")
		kendra_updateQuerySuggestionsBlockListCmd.Flags().String("name", "", "A new name for the block list.")
		kendra_updateQuerySuggestionsBlockListCmd.Flags().String("role-arn", "", "The IAM (Identity and Access Management) role used to access the block list text file in S3.")
		kendra_updateQuerySuggestionsBlockListCmd.Flags().String("source-s3-path", "", "The S3 path where your block list text file sits in S3.")
		kendra_updateQuerySuggestionsBlockListCmd.MarkFlagRequired("id")
		kendra_updateQuerySuggestionsBlockListCmd.MarkFlagRequired("index-id")
	})
	kendraCmd.AddCommand(kendra_updateQuerySuggestionsBlockListCmd)
}
