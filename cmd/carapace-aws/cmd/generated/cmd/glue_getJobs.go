package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getJobsCmd = &cobra.Command{
	Use:   "get-jobs",
	Short: "Retrieves all current job definitions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getJobsCmd).Standalone()

		glue_getJobsCmd.Flags().String("max-results", "", "The maximum size of the response.")
		glue_getJobsCmd.Flags().String("next-token", "", "A continuation token, if this is a continuation call.")
	})
	glueCmd.AddCommand(glue_getJobsCmd)
}
