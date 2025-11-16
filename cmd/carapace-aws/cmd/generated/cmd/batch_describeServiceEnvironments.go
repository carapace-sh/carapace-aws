package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_describeServiceEnvironmentsCmd = &cobra.Command{
	Use:   "describe-service-environments",
	Short: "Describes one or more of your service environments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_describeServiceEnvironmentsCmd).Standalone()

	batch_describeServiceEnvironmentsCmd.Flags().String("max-results", "", "The maximum number of results returned by `DescribeServiceEnvironments` in paginated output.")
	batch_describeServiceEnvironmentsCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated `DescribeServiceEnvironments` request where `maxResults` was used and the results exceeded the value of that parameter.")
	batch_describeServiceEnvironmentsCmd.Flags().String("service-environments", "", "An array of service environment names or ARN entries.")
	batchCmd.AddCommand(batch_describeServiceEnvironmentsCmd)
}
