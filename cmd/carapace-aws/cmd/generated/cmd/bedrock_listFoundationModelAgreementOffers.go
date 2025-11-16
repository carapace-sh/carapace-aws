package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_listFoundationModelAgreementOffersCmd = &cobra.Command{
	Use:   "list-foundation-model-agreement-offers",
	Short: "Get the offers associated with the specified model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_listFoundationModelAgreementOffersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_listFoundationModelAgreementOffersCmd).Standalone()

		bedrock_listFoundationModelAgreementOffersCmd.Flags().String("model-id", "", "Model Id of the foundation model.")
		bedrock_listFoundationModelAgreementOffersCmd.Flags().String("offer-type", "", "Type of offer associated with the model.")
		bedrock_listFoundationModelAgreementOffersCmd.MarkFlagRequired("model-id")
	})
	bedrockCmd.AddCommand(bedrock_listFoundationModelAgreementOffersCmd)
}
