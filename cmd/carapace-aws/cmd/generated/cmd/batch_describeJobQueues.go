package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_describeJobQueuesCmd = &cobra.Command{
	Use:   "describe-job-queues",
	Short: "Describes one or more of your job queues.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_describeJobQueuesCmd).Standalone()

	batch_describeJobQueuesCmd.Flags().String("job-queues", "", "A list of up to 100 queue names or full queue Amazon Resource Name (ARN) entries.")
	batch_describeJobQueuesCmd.Flags().String("max-results", "", "The maximum number of results returned by `DescribeJobQueues` in paginated output.")
	batch_describeJobQueuesCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated `DescribeJobQueues` request where `maxResults` was used and the results exceeded the value of that parameter.")
	batchCmd.AddCommand(batch_describeJobQueuesCmd)
}
