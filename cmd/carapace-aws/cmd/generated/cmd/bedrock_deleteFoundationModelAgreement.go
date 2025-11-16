package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_deleteFoundationModelAgreementCmd = &cobra.Command{
	Use:   "delete-foundation-model-agreement",
	Short: "Delete the model access agreement for the specified model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_deleteFoundationModelAgreementCmd).Standalone()

	bedrock_deleteFoundationModelAgreementCmd.Flags().String("model-id", "", "Model Id of the model access to delete.")
	bedrock_deleteFoundationModelAgreementCmd.MarkFlagRequired("model-id")
	bedrockCmd.AddCommand(bedrock_deleteFoundationModelAgreementCmd)
}
