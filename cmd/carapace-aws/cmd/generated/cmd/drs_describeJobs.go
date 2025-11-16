package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_describeJobsCmd = &cobra.Command{
	Use:   "describe-jobs",
	Short: "Returns a list of Jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_describeJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(drs_describeJobsCmd).Standalone()

		drs_describeJobsCmd.Flags().String("filters", "", "A set of filters by which to return Jobs.")
		drs_describeJobsCmd.Flags().String("max-results", "", "Maximum number of Jobs to retrieve.")
		drs_describeJobsCmd.Flags().String("next-token", "", "The token of the next Job to retrieve.")
	})
	drsCmd.AddCommand(drs_describeJobsCmd)
}
