package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_listLoaderJobsCmd = &cobra.Command{
	Use:   "list-loader-jobs",
	Short: "Retrieves a list of the `loadIds` for all active loader jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_listLoaderJobsCmd).Standalone()

	neptunedata_listLoaderJobsCmd.Flags().Bool("include-queued-loads", false, "An optional parameter that can be used to exclude the load IDs of queued load requests when requesting a list of load IDs by setting the parameter to `FALSE`.")
	neptunedata_listLoaderJobsCmd.Flags().String("limit", "", "The number of load IDs to list.")
	neptunedata_listLoaderJobsCmd.Flags().Bool("no-include-queued-loads", false, "An optional parameter that can be used to exclude the load IDs of queued load requests when requesting a list of load IDs by setting the parameter to `FALSE`.")
	neptunedata_listLoaderJobsCmd.Flag("no-include-queued-loads").Hidden = true
	neptunedataCmd.AddCommand(neptunedata_listLoaderJobsCmd)
}
