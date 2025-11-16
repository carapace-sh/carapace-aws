package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_startJobRunCmd = &cobra.Command{
	Use:   "start-job-run",
	Short: "Runs a DataBrew job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_startJobRunCmd).Standalone()

	databrew_startJobRunCmd.Flags().String("name", "", "The name of the job to be run.")
	databrew_startJobRunCmd.MarkFlagRequired("name")
	databrewCmd.AddCommand(databrew_startJobRunCmd)
}
