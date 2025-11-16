package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getTriggersCmd = &cobra.Command{
	Use:   "get-triggers",
	Short: "Gets all the triggers associated with a job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getTriggersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getTriggersCmd).Standalone()

		glue_getTriggersCmd.Flags().String("dependent-job-name", "", "The name of the job to retrieve triggers for.")
		glue_getTriggersCmd.Flags().String("max-results", "", "The maximum size of the response.")
		glue_getTriggersCmd.Flags().String("next-token", "", "A continuation token, if this is a continuation call.")
	})
	glueCmd.AddCommand(glue_getTriggersCmd)
}
