package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_getFileUploadUrlCmd = &cobra.Command{
	Use:   "get-file-upload-url",
	Short: "The `GetFileUploadURL` operation generates and returns a temporary URL.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_getFileUploadUrlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mturk_getFileUploadUrlCmd).Standalone()

		mturk_getFileUploadUrlCmd.Flags().String("assignment-id", "", "The ID of the assignment that contains the question with a FileUploadAnswer.")
		mturk_getFileUploadUrlCmd.Flags().String("question-identifier", "", "The identifier of the question with a FileUploadAnswer, as specified in the QuestionForm of the HIT.")
		mturk_getFileUploadUrlCmd.MarkFlagRequired("assignment-id")
		mturk_getFileUploadUrlCmd.MarkFlagRequired("question-identifier")
	})
	mturkCmd.AddCommand(mturk_getFileUploadUrlCmd)
}
