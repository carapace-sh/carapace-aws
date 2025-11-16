package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowball_listJobsCmd = &cobra.Command{
	Use:   "list-jobs",
	Short: "Returns an array of `JobListEntry` objects of the specified length.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowball_listJobsCmd).Standalone()

	snowball_listJobsCmd.Flags().String("max-results", "", "The number of `JobListEntry` objects to return.")
	snowball_listJobsCmd.Flags().String("next-token", "", "HTTP requests are stateless.")
	snowballCmd.AddCommand(snowball_listJobsCmd)
}
