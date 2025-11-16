package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzer_startResourceScanCmd = &cobra.Command{
	Use:   "start-resource-scan",
	Short: "Immediately starts a scan of the policies applied to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzer_startResourceScanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(accessanalyzer_startResourceScanCmd).Standalone()

		accessanalyzer_startResourceScanCmd.Flags().String("analyzer-arn", "", "The [ARN of the analyzer](https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-getting-started.html#permission-resources) to use to scan the policies applied to the specified resource.")
		accessanalyzer_startResourceScanCmd.Flags().String("resource-arn", "", "The ARN of the resource to scan.")
		accessanalyzer_startResourceScanCmd.Flags().String("resource-owner-account", "", "The Amazon Web Services account ID that owns the resource.")
		accessanalyzer_startResourceScanCmd.MarkFlagRequired("analyzer-arn")
		accessanalyzer_startResourceScanCmd.MarkFlagRequired("resource-arn")
	})
	accessanalyzerCmd.AddCommand(accessanalyzer_startResourceScanCmd)
}
