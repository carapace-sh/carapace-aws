package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvert_searchJobsCmd = &cobra.Command{
	Use:   "search-jobs",
	Short: "Retrieve a JSON array that includes job details for up to twenty of your most recent jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvert_searchJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconvert_searchJobsCmd).Standalone()

		mediaconvert_searchJobsCmd.Flags().String("input-file", "", "Optional.")
		mediaconvert_searchJobsCmd.Flags().String("max-results", "", "Optional.")
		mediaconvert_searchJobsCmd.Flags().String("next-token", "", "Optional.")
		mediaconvert_searchJobsCmd.Flags().String("order", "", "Optional.")
		mediaconvert_searchJobsCmd.Flags().String("queue", "", "Optional.")
		mediaconvert_searchJobsCmd.Flags().String("status", "", "Optional.")
	})
	mediaconvertCmd.AddCommand(mediaconvert_searchJobsCmd)
}
