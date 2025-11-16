package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_deleteProvisionedModelThroughputCmd = &cobra.Command{
	Use:   "delete-provisioned-model-throughput",
	Short: "Deletes a Provisioned Throughput.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_deleteProvisionedModelThroughputCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_deleteProvisionedModelThroughputCmd).Standalone()

		bedrock_deleteProvisionedModelThroughputCmd.Flags().String("provisioned-model-id", "", "The Amazon Resource Name (ARN) or name of the Provisioned Throughput.")
		bedrock_deleteProvisionedModelThroughputCmd.MarkFlagRequired("provisioned-model-id")
	})
	bedrockCmd.AddCommand(bedrock_deleteProvisionedModelThroughputCmd)
}
