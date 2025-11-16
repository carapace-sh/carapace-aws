package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_listFiltersCmd = &cobra.Command{
	Use:   "list-filters",
	Short: "Lists the filters associated with your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_listFiltersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_listFiltersCmd).Standalone()

		inspector2_listFiltersCmd.Flags().String("action", "", "The action the filter applies to matched findings.")
		inspector2_listFiltersCmd.Flags().String("arns", "", "The Amazon resource number (ARN) of the filter.")
		inspector2_listFiltersCmd.Flags().String("max-results", "", "The maximum number of results the response can return.")
		inspector2_listFiltersCmd.Flags().String("next-token", "", "A token to use for paginating results that are returned in the response.")
	})
	inspector2Cmd.AddCommand(inspector2_listFiltersCmd)
}
