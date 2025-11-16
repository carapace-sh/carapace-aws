package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_describeJobDefinitionsCmd = &cobra.Command{
	Use:   "describe-job-definitions",
	Short: "Describes a list of job definitions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_describeJobDefinitionsCmd).Standalone()

	batch_describeJobDefinitionsCmd.Flags().String("job-definition-name", "", "The name of the job definition to describe.")
	batch_describeJobDefinitionsCmd.Flags().String("job-definitions", "", "A list of up to 100 job definitions.")
	batch_describeJobDefinitionsCmd.Flags().String("max-results", "", "The maximum number of results returned by `DescribeJobDefinitions` in paginated output.")
	batch_describeJobDefinitionsCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated `DescribeJobDefinitions` request where `maxResults` was used and the results exceeded the value of that parameter.")
	batch_describeJobDefinitionsCmd.Flags().String("status", "", "The status used to filter job definitions.")
	batchCmd.AddCommand(batch_describeJobDefinitionsCmd)
}
