package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_getProvisionedModelThroughputCmd = &cobra.Command{
	Use:   "get-provisioned-model-throughput",
	Short: "Returns details for a Provisioned Throughput.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_getProvisionedModelThroughputCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_getProvisionedModelThroughputCmd).Standalone()

		bedrock_getProvisionedModelThroughputCmd.Flags().String("provisioned-model-id", "", "The Amazon Resource Name (ARN) or name of the Provisioned Throughput.")
		bedrock_getProvisionedModelThroughputCmd.MarkFlagRequired("provisioned-model-id")
	})
	bedrockCmd.AddCommand(bedrock_getProvisionedModelThroughputCmd)
}
