package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_getFindingsCmd = &cobra.Command{
	Use:   "get-findings",
	Short: "Returns a list of findings that match the specified criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_getFindingsCmd).Standalone()

	securityhub_getFindingsCmd.Flags().String("filters", "", "The finding attributes used to define a condition to filter the returned findings.")
	securityhub_getFindingsCmd.Flags().String("max-results", "", "The maximum number of findings to return.")
	securityhub_getFindingsCmd.Flags().String("next-token", "", "The token that is required for pagination.")
	securityhub_getFindingsCmd.Flags().String("sort-criteria", "", "The finding attributes used to sort the list of returned findings.")
	securityhubCmd.AddCommand(securityhub_getFindingsCmd)
}
