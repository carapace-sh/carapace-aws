package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_getSensitiveDataOccurrencesAvailabilityCmd = &cobra.Command{
	Use:   "get-sensitive-data-occurrences-availability",
	Short: "Checks whether occurrences of sensitive data can be retrieved for a finding.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_getSensitiveDataOccurrencesAvailabilityCmd).Standalone()

	macie2_getSensitiveDataOccurrencesAvailabilityCmd.Flags().String("finding-id", "", "The unique identifier for the finding.")
	macie2_getSensitiveDataOccurrencesAvailabilityCmd.MarkFlagRequired("finding-id")
	macie2Cmd.AddCommand(macie2_getSensitiveDataOccurrencesAvailabilityCmd)
}
