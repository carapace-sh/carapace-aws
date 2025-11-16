package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var marketplaceAgreement_getAgreementTermsCmd = &cobra.Command{
	Use:   "get-agreement-terms",
	Short: "Obtains details about the terms in an agreement that you participated in as proposer or acceptor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(marketplaceAgreement_getAgreementTermsCmd).Standalone()

	marketplaceAgreement_getAgreementTermsCmd.Flags().String("agreement-id", "", "The unique identifier of the agreement.")
	marketplaceAgreement_getAgreementTermsCmd.Flags().String("max-results", "", "The maximum number of agreements to return in the response.")
	marketplaceAgreement_getAgreementTermsCmd.Flags().String("next-token", "", "A token to specify where to start pagination")
	marketplaceAgreement_getAgreementTermsCmd.MarkFlagRequired("agreement-id")
	marketplaceAgreementCmd.AddCommand(marketplaceAgreement_getAgreementTermsCmd)
}
