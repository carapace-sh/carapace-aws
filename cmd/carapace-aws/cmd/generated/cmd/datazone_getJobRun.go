package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getJobRunCmd = &cobra.Command{
	Use:   "get-job-run",
	Short: "The details of the job run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getJobRunCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_getJobRunCmd).Standalone()

		datazone_getJobRunCmd.Flags().String("domain-identifier", "", "The ID of the domain.")
		datazone_getJobRunCmd.Flags().String("identifier", "", "The ID of the job run.")
		datazone_getJobRunCmd.MarkFlagRequired("domain-identifier")
		datazone_getJobRunCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_getJobRunCmd)
}
