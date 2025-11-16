package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var trustedadvisor_listChecksCmd = &cobra.Command{
	Use:   "list-checks",
	Short: "List a filterable set of Checks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trustedadvisor_listChecksCmd).Standalone()

	trustedadvisor_listChecksCmd.Flags().String("aws-service", "", "The aws service associated with the check")
	trustedadvisor_listChecksCmd.Flags().String("language", "", "The ISO 639-1 code for the language that you want your checks to appear in.")
	trustedadvisor_listChecksCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	trustedadvisor_listChecksCmd.Flags().String("next-token", "", "The token for the next set of results.")
	trustedadvisor_listChecksCmd.Flags().String("pillar", "", "The pillar of the check")
	trustedadvisor_listChecksCmd.Flags().String("source", "", "The source of the check")
	trustedadvisorCmd.AddCommand(trustedadvisor_listChecksCmd)
}
