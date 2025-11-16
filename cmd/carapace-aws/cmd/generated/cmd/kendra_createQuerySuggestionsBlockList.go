package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_createQuerySuggestionsBlockListCmd = &cobra.Command{
	Use:   "create-query-suggestions-block-list",
	Short: "Creates a block list to exlcude certain queries from suggestions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_createQuerySuggestionsBlockListCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendra_createQuerySuggestionsBlockListCmd).Standalone()

		kendra_createQuerySuggestionsBlockListCmd.Flags().String("client-token", "", "A token that you provide to identify the request to create a query suggestions block list.")
		kendra_createQuerySuggestionsBlockListCmd.Flags().String("description", "", "A description for the block list.")
		kendra_createQuerySuggestionsBlockListCmd.Flags().String("index-id", "", "The identifier of the index you want to create a query suggestions block list for.")
		kendra_createQuerySuggestionsBlockListCmd.Flags().String("name", "", "A name for the block list.")
		kendra_createQuerySuggestionsBlockListCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role with permission to access your S3 bucket that contains the block list text file.")
		kendra_createQuerySuggestionsBlockListCmd.Flags().String("source-s3-path", "", "The S3 path to your block list text file in your S3 bucket.")
		kendra_createQuerySuggestionsBlockListCmd.Flags().String("tags", "", "A list of key-value pairs that identify or categorize the block list.")
		kendra_createQuerySuggestionsBlockListCmd.MarkFlagRequired("index-id")
		kendra_createQuerySuggestionsBlockListCmd.MarkFlagRequired("name")
		kendra_createQuerySuggestionsBlockListCmd.MarkFlagRequired("role-arn")
		kendra_createQuerySuggestionsBlockListCmd.MarkFlagRequired("source-s3-path")
	})
	kendraCmd.AddCommand(kendra_createQuerySuggestionsBlockListCmd)
}
