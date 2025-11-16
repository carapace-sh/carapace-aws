package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_updateThesaurusCmd = &cobra.Command{
	Use:   "update-thesaurus",
	Short: "Updates a thesaurus for an index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_updateThesaurusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendra_updateThesaurusCmd).Standalone()

		kendra_updateThesaurusCmd.Flags().String("description", "", "A new description for the thesaurus.")
		kendra_updateThesaurusCmd.Flags().String("id", "", "The identifier of the thesaurus you want to update.")
		kendra_updateThesaurusCmd.Flags().String("index-id", "", "The identifier of the index for the thesaurus.")
		kendra_updateThesaurusCmd.Flags().String("name", "", "A new name for the thesaurus.")
		kendra_updateThesaurusCmd.Flags().String("role-arn", "", "An IAM role that gives Amazon Kendra permissions to access thesaurus file specified in `SourceS3Path`.")
		kendra_updateThesaurusCmd.Flags().String("source-s3-path", "", "")
		kendra_updateThesaurusCmd.MarkFlagRequired("id")
		kendra_updateThesaurusCmd.MarkFlagRequired("index-id")
	})
	kendraCmd.AddCommand(kendra_updateThesaurusCmd)
}
