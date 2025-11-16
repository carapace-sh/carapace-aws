package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_listProvisionedModelThroughputsCmd = &cobra.Command{
	Use:   "list-provisioned-model-throughputs",
	Short: "Lists the Provisioned Throughputs in the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_listProvisionedModelThroughputsCmd).Standalone()

	bedrock_listProvisionedModelThroughputsCmd.Flags().String("creation-time-after", "", "A filter that returns Provisioned Throughputs created after the specified time.")
	bedrock_listProvisionedModelThroughputsCmd.Flags().String("creation-time-before", "", "A filter that returns Provisioned Throughputs created before the specified time.")
	bedrock_listProvisionedModelThroughputsCmd.Flags().String("max-results", "", "THe maximum number of results to return in the response.")
	bedrock_listProvisionedModelThroughputsCmd.Flags().String("model-arn-equals", "", "A filter that returns Provisioned Throughputs whose model Amazon Resource Name (ARN) is equal to the value that you specify.")
	bedrock_listProvisionedModelThroughputsCmd.Flags().String("name-contains", "", "A filter that returns Provisioned Throughputs if their name contains the expression that you specify.")
	bedrock_listProvisionedModelThroughputsCmd.Flags().String("next-token", "", "If there are more results than the number you specified in the `maxResults` field, the response returns a `nextToken` value.")
	bedrock_listProvisionedModelThroughputsCmd.Flags().String("sort-by", "", "The field by which to sort the returned list of Provisioned Throughputs.")
	bedrock_listProvisionedModelThroughputsCmd.Flags().String("sort-order", "", "The sort order of the results.")
	bedrock_listProvisionedModelThroughputsCmd.Flags().String("status-equals", "", "A filter that returns Provisioned Throughputs if their statuses matches the value that you specify.")
	bedrockCmd.AddCommand(bedrock_listProvisionedModelThroughputsCmd)
}
