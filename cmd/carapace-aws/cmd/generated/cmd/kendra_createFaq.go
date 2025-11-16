package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_createFaqCmd = &cobra.Command{
	Use:   "create-faq",
	Short: "Creates a set of frequently ask questions (FAQs) using a specified FAQ file stored in an Amazon S3 bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_createFaqCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendra_createFaqCmd).Standalone()

		kendra_createFaqCmd.Flags().String("client-token", "", "A token that you provide to identify the request to create a FAQ.")
		kendra_createFaqCmd.Flags().String("description", "", "A description for the FAQ.")
		kendra_createFaqCmd.Flags().String("file-format", "", "The format of the FAQ input file.")
		kendra_createFaqCmd.Flags().String("index-id", "", "The identifier of the index for the FAQ.")
		kendra_createFaqCmd.Flags().String("language-code", "", "The code for a language.")
		kendra_createFaqCmd.Flags().String("name", "", "A name for the FAQ.")
		kendra_createFaqCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role with permission to access the S3 bucket that contains the FAQ file.")
		kendra_createFaqCmd.Flags().String("s3-path", "", "The path to the FAQ file in S3.")
		kendra_createFaqCmd.Flags().String("tags", "", "A list of key-value pairs that identify the FAQ.")
		kendra_createFaqCmd.MarkFlagRequired("index-id")
		kendra_createFaqCmd.MarkFlagRequired("name")
		kendra_createFaqCmd.MarkFlagRequired("role-arn")
		kendra_createFaqCmd.MarkFlagRequired("s3-path")
	})
	kendraCmd.AddCommand(kendra_createFaqCmd)
}
