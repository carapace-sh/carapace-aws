package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_getImportedModelCmd = &cobra.Command{
	Use:   "get-imported-model",
	Short: "Gets properties associated with a customized model you imported.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_getImportedModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_getImportedModelCmd).Standalone()

		bedrock_getImportedModelCmd.Flags().String("model-identifier", "", "Name or Amazon Resource Name (ARN) of the imported model.")
		bedrock_getImportedModelCmd.MarkFlagRequired("model-identifier")
	})
	bedrockCmd.AddCommand(bedrock_getImportedModelCmd)
}
