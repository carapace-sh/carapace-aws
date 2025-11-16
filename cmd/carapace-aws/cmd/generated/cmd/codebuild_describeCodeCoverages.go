package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_describeCodeCoveragesCmd = &cobra.Command{
	Use:   "describe-code-coverages",
	Short: "Retrieves one or more code coverage reports.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_describeCodeCoveragesCmd).Standalone()

	codebuild_describeCodeCoveragesCmd.Flags().String("max-line-coverage-percentage", "", "The maximum line coverage percentage to report.")
	codebuild_describeCodeCoveragesCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	codebuild_describeCodeCoveragesCmd.Flags().String("min-line-coverage-percentage", "", "The minimum line coverage percentage to report.")
	codebuild_describeCodeCoveragesCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous call to `DescribeCodeCoverages`.")
	codebuild_describeCodeCoveragesCmd.Flags().String("report-arn", "", "The ARN of the report for which test cases are returned.")
	codebuild_describeCodeCoveragesCmd.Flags().String("sort-by", "", "Specifies how the results are sorted.")
	codebuild_describeCodeCoveragesCmd.Flags().String("sort-order", "", "Specifies if the results are sorted in ascending or descending order.")
	codebuild_describeCodeCoveragesCmd.MarkFlagRequired("report-arn")
	codebuildCmd.AddCommand(codebuild_describeCodeCoveragesCmd)
}
