package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var artifact_listCustomerAgreementsCmd = &cobra.Command{
	Use:   "list-customer-agreements",
	Short: "List active customer-agreements applicable to calling identity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(artifact_listCustomerAgreementsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(artifact_listCustomerAgreementsCmd).Standalone()

		artifact_listCustomerAgreementsCmd.Flags().String("max-results", "", "Maximum number of resources to return in the paginated response.")
		artifact_listCustomerAgreementsCmd.Flags().String("next-token", "", "Pagination token to request the next page of resources.")
	})
	artifactCmd.AddCommand(artifact_listCustomerAgreementsCmd)
}
