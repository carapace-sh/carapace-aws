package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elastictranscoder_testRoleCmd = &cobra.Command{
	Use:   "test-role",
	Short: "The TestRole operation tests the IAM role used to create the pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elastictranscoder_testRoleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elastictranscoder_testRoleCmd).Standalone()

		elastictranscoder_testRoleCmd.Flags().String("input-bucket", "", "The Amazon S3 bucket that contains media files to be transcoded.")
		elastictranscoder_testRoleCmd.Flags().String("output-bucket", "", "The Amazon S3 bucket that Elastic Transcoder writes transcoded media files to.")
		elastictranscoder_testRoleCmd.Flags().String("role", "", "The IAM Amazon Resource Name (ARN) for the role that you want Elastic Transcoder to test.")
		elastictranscoder_testRoleCmd.Flags().String("topics", "", "The ARNs of one or more Amazon Simple Notification Service (Amazon SNS) topics that you want the action to send a test notification to.")
		elastictranscoder_testRoleCmd.MarkFlagRequired("input-bucket")
		elastictranscoder_testRoleCmd.MarkFlagRequired("output-bucket")
		elastictranscoder_testRoleCmd.MarkFlagRequired("role")
		elastictranscoder_testRoleCmd.MarkFlagRequired("topics")
	})
	elastictranscoderCmd.AddCommand(elastictranscoder_testRoleCmd)
}
