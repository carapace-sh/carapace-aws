package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_createFoundationModelAgreementCmd = &cobra.Command{
	Use:   "create-foundation-model-agreement",
	Short: "Request a model access agreement for the specified model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_createFoundationModelAgreementCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_createFoundationModelAgreementCmd).Standalone()

		bedrock_createFoundationModelAgreementCmd.Flags().String("model-id", "", "Model Id of the model for the access request.")
		bedrock_createFoundationModelAgreementCmd.Flags().String("offer-token", "", "An offer token encapsulates the information for an offer.")
		bedrock_createFoundationModelAgreementCmd.MarkFlagRequired("model-id")
		bedrock_createFoundationModelAgreementCmd.MarkFlagRequired("offer-token")
	})
	bedrockCmd.AddCommand(bedrock_createFoundationModelAgreementCmd)
}
