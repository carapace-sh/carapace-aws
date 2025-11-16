package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_describeJobRunCmd = &cobra.Command{
	Use:   "describe-job-run",
	Short: "Represents one run of a DataBrew job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_describeJobRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(databrew_describeJobRunCmd).Standalone()

		databrew_describeJobRunCmd.Flags().String("name", "", "The name of the job being processed during this run.")
		databrew_describeJobRunCmd.Flags().String("run-id", "", "The unique identifier of the job run.")
		databrew_describeJobRunCmd.MarkFlagRequired("name")
		databrew_describeJobRunCmd.MarkFlagRequired("run-id")
	})
	databrewCmd.AddCommand(databrew_describeJobRunCmd)
}
