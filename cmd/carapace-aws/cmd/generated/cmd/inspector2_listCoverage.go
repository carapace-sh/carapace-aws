package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_listCoverageCmd = &cobra.Command{
	Use:   "list-coverage",
	Short: "Lists coverage details for your environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_listCoverageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_listCoverageCmd).Standalone()

		inspector2_listCoverageCmd.Flags().String("filter-criteria", "", "An object that contains details on the filters to apply to the coverage data for your environment.")
		inspector2_listCoverageCmd.Flags().String("max-results", "", "The maximum number of results the response can return.")
		inspector2_listCoverageCmd.Flags().String("next-token", "", "A token to use for paginating results that are returned in the response.")
	})
	inspector2Cmd.AddCommand(inspector2_listCoverageCmd)
}
