package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_listCustomModelDeploymentsCmd = &cobra.Command{
	Use:   "list-custom-model-deployments",
	Short: "Lists custom model deployments in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_listCustomModelDeploymentsCmd).Standalone()

	bedrock_listCustomModelDeploymentsCmd.Flags().String("created-after", "", "Filters deployments created after the specified date and time.")
	bedrock_listCustomModelDeploymentsCmd.Flags().String("created-before", "", "Filters deployments created before the specified date and time.")
	bedrock_listCustomModelDeploymentsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	bedrock_listCustomModelDeploymentsCmd.Flags().String("model-arn-equals", "", "Filters deployments by the Amazon Resource Name (ARN) of the associated custom model.")
	bedrock_listCustomModelDeploymentsCmd.Flags().String("name-contains", "", "Filters deployments whose names contain the specified string.")
	bedrock_listCustomModelDeploymentsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	bedrock_listCustomModelDeploymentsCmd.Flags().String("sort-by", "", "The field to sort the results by.")
	bedrock_listCustomModelDeploymentsCmd.Flags().String("sort-order", "", "The sort order for the results.")
	bedrock_listCustomModelDeploymentsCmd.Flags().String("status-equals", "", "Filters deployments by status.")
	bedrockCmd.AddCommand(bedrock_listCustomModelDeploymentsCmd)
}
