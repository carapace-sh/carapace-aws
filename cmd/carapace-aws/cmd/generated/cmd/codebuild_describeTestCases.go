package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_describeTestCasesCmd = &cobra.Command{
	Use:   "describe-test-cases",
	Short: "Returns a list of details about test cases for a report.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_describeTestCasesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codebuild_describeTestCasesCmd).Standalone()

		codebuild_describeTestCasesCmd.Flags().String("filter", "", "A `TestCaseFilter` object used to filter the returned reports.")
		codebuild_describeTestCasesCmd.Flags().String("max-results", "", "The maximum number of paginated test cases returned per response.")
		codebuild_describeTestCasesCmd.Flags().String("next-token", "", "During a previous call, the maximum number of items that can be returned is the value specified in `maxResults`.")
		codebuild_describeTestCasesCmd.Flags().String("report-arn", "", "The ARN of the report for which test cases are returned.")
		codebuild_describeTestCasesCmd.MarkFlagRequired("report-arn")
	})
	codebuildCmd.AddCommand(codebuild_describeTestCasesCmd)
}
