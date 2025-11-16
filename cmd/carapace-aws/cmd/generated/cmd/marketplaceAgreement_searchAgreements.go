package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var marketplaceAgreement_searchAgreementsCmd = &cobra.Command{
	Use:   "search-agreements",
	Short: "Searches across all agreements that a proposer or an acceptor has in AWS Marketplace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(marketplaceAgreement_searchAgreementsCmd).Standalone()

	marketplaceAgreement_searchAgreementsCmd.Flags().String("catalog", "", "The catalog in which the agreement was created.")
	marketplaceAgreement_searchAgreementsCmd.Flags().String("filters", "", "The filter name and value pair used to return a specific list of results.")
	marketplaceAgreement_searchAgreementsCmd.Flags().String("max-results", "", "The maximum number of agreements to return in the response.")
	marketplaceAgreement_searchAgreementsCmd.Flags().String("next-token", "", "A token to specify where to start pagination.")
	marketplaceAgreement_searchAgreementsCmd.Flags().String("sort", "", "An object that contains the `SortBy` and `SortOrder` attributes.")
	marketplaceAgreementCmd.AddCommand(marketplaceAgreement_searchAgreementsCmd)
}
