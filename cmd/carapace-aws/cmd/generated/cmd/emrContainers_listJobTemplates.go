package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrContainers_listJobTemplatesCmd = &cobra.Command{
	Use:   "list-job-templates",
	Short: "Lists job templates based on a set of parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrContainers_listJobTemplatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emrContainers_listJobTemplatesCmd).Standalone()

		emrContainers_listJobTemplatesCmd.Flags().String("created-after", "", "The date and time after which the job templates were created.")
		emrContainers_listJobTemplatesCmd.Flags().String("created-before", "", "The date and time before which the job templates were created.")
		emrContainers_listJobTemplatesCmd.Flags().String("max-results", "", "The maximum number of job templates that can be listed.")
		emrContainers_listJobTemplatesCmd.Flags().String("next-token", "", "The token for the next set of job templates to return.")
	})
	emrContainersCmd.AddCommand(emrContainers_listJobTemplatesCmd)
}
