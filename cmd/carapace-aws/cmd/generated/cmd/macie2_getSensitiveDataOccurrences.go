package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_getSensitiveDataOccurrencesCmd = &cobra.Command{
	Use:   "get-sensitive-data-occurrences",
	Short: "Retrieves occurrences of sensitive data reported by a finding.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_getSensitiveDataOccurrencesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_getSensitiveDataOccurrencesCmd).Standalone()

		macie2_getSensitiveDataOccurrencesCmd.Flags().String("finding-id", "", "The unique identifier for the finding.")
		macie2_getSensitiveDataOccurrencesCmd.MarkFlagRequired("finding-id")
	})
	macie2Cmd.AddCommand(macie2_getSensitiveDataOccurrencesCmd)
}
