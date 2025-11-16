package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_createThesaurusCmd = &cobra.Command{
	Use:   "create-thesaurus",
	Short: "Creates a thesaurus for an index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_createThesaurusCmd).Standalone()

	kendra_createThesaurusCmd.Flags().String("client-token", "", "A token that you provide to identify the request to create a thesaurus.")
	kendra_createThesaurusCmd.Flags().String("description", "", "A description for the thesaurus.")
	kendra_createThesaurusCmd.Flags().String("index-id", "", "The identifier of the index for the thesaurus.")
	kendra_createThesaurusCmd.Flags().String("name", "", "A name for the thesaurus.")
	kendra_createThesaurusCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role with permission to access your S3 bucket that contains the thesaurus file.")
	kendra_createThesaurusCmd.Flags().String("source-s3-path", "", "The path to the thesaurus file in S3.")
	kendra_createThesaurusCmd.Flags().String("tags", "", "A list of key-value pairs that identify or categorize the thesaurus.")
	kendra_createThesaurusCmd.MarkFlagRequired("index-id")
	kendra_createThesaurusCmd.MarkFlagRequired("name")
	kendra_createThesaurusCmd.MarkFlagRequired("role-arn")
	kendra_createThesaurusCmd.MarkFlagRequired("source-s3-path")
	kendraCmd.AddCommand(kendra_createThesaurusCmd)
}
