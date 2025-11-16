package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_describeComputeEnvironmentsCmd = &cobra.Command{
	Use:   "describe-compute-environments",
	Short: "Describes one or more of your compute environments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_describeComputeEnvironmentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(batch_describeComputeEnvironmentsCmd).Standalone()

		batch_describeComputeEnvironmentsCmd.Flags().String("compute-environments", "", "A list of up to 100 compute environment names or full Amazon Resource Name (ARN) entries.")
		batch_describeComputeEnvironmentsCmd.Flags().String("max-results", "", "The maximum number of cluster results returned by `DescribeComputeEnvironments` in paginated output.")
		batch_describeComputeEnvironmentsCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated `DescribeComputeEnvironments` request where `maxResults` was used and the results exceeded the value of that parameter.")
	})
	batchCmd.AddCommand(batch_describeComputeEnvironmentsCmd)
}
