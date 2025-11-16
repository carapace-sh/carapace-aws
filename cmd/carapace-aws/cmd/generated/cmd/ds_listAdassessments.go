package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_listAdassessmentsCmd = &cobra.Command{
	Use:   "list-adassessments",
	Short: "Retrieves a list of directory assessments for the specified directory or all assessments in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_listAdassessmentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_listAdassessmentsCmd).Standalone()

		ds_listAdassessmentsCmd.Flags().String("directory-id", "", "The identifier of the directory for which to list assessments.")
		ds_listAdassessmentsCmd.Flags().String("limit", "", "The maximum number of assessment summaries to return.")
		ds_listAdassessmentsCmd.Flags().String("next-token", "", "The pagination token from a previous request to [ListADAssessments]().")
	})
	dsCmd.AddCommand(ds_listAdassessmentsCmd)
}
