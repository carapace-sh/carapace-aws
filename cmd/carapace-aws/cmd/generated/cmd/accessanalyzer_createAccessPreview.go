package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzer_createAccessPreviewCmd = &cobra.Command{
	Use:   "create-access-preview",
	Short: "Creates an access preview that allows you to preview IAM Access Analyzer findings for your resource before deploying resource permissions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzer_createAccessPreviewCmd).Standalone()

	accessanalyzer_createAccessPreviewCmd.Flags().String("analyzer-arn", "", "The [ARN of the account analyzer](https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-getting-started.html#permission-resources) used to generate the access preview.")
	accessanalyzer_createAccessPreviewCmd.Flags().String("client-token", "", "A client token.")
	accessanalyzer_createAccessPreviewCmd.Flags().String("configurations", "", "Access control configuration for your resource that is used to generate the access preview.")
	accessanalyzer_createAccessPreviewCmd.MarkFlagRequired("analyzer-arn")
	accessanalyzer_createAccessPreviewCmd.MarkFlagRequired("configurations")
	accessanalyzerCmd.AddCommand(accessanalyzer_createAccessPreviewCmd)
}
