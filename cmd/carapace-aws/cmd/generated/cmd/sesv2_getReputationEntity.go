package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_getReputationEntityCmd = &cobra.Command{
	Use:   "get-reputation-entity",
	Short: "Retrieve information about a specific reputation entity, including its reputation management policy, customer-managed status, Amazon Web Services Amazon SES-managed status, and aggregate sending status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_getReputationEntityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_getReputationEntityCmd).Standalone()

		sesv2_getReputationEntityCmd.Flags().String("reputation-entity-reference", "", "The unique identifier for the reputation entity.")
		sesv2_getReputationEntityCmd.Flags().String("reputation-entity-type", "", "The type of reputation entity.")
		sesv2_getReputationEntityCmd.MarkFlagRequired("reputation-entity-reference")
		sesv2_getReputationEntityCmd.MarkFlagRequired("reputation-entity-type")
	})
	sesv2Cmd.AddCommand(sesv2_getReputationEntityCmd)
}
