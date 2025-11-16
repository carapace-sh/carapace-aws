package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var marketplaceAgreement_describeAgreementCmd = &cobra.Command{
	Use:   "describe-agreement",
	Short: "Provides details about an agreement, such as the proposer, acceptor, start date, and end date.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(marketplaceAgreement_describeAgreementCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(marketplaceAgreement_describeAgreementCmd).Standalone()

		marketplaceAgreement_describeAgreementCmd.Flags().String("agreement-id", "", "The unique identifier of the agreement.")
		marketplaceAgreement_describeAgreementCmd.MarkFlagRequired("agreement-id")
	})
	marketplaceAgreementCmd.AddCommand(marketplaceAgreement_describeAgreementCmd)
}
