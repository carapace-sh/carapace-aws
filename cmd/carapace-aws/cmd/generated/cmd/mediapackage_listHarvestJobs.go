package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackage_listHarvestJobsCmd = &cobra.Command{
	Use:   "list-harvest-jobs",
	Short: "Returns a collection of HarvestJob records.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackage_listHarvestJobsCmd).Standalone()

	mediapackage_listHarvestJobsCmd.Flags().String("include-channel-id", "", "When specified, the request will return only HarvestJobs associated with the given Channel ID.")
	mediapackage_listHarvestJobsCmd.Flags().String("include-status", "", "When specified, the request will return only HarvestJobs in the given status.")
	mediapackage_listHarvestJobsCmd.Flags().String("max-results", "", "The upper bound on the number of records to return.")
	mediapackage_listHarvestJobsCmd.Flags().String("next-token", "", "A token used to resume pagination from the end of a previous request.")
	mediapackageCmd.AddCommand(mediapackage_listHarvestJobsCmd)
}
