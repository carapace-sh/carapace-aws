package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var importexport_getStatusCmd = &cobra.Command{
	Use:   "get-status",
	Short: "This operation returns information about a job, including where the job is in the processing pipeline, the status of the results, and the signature value associated with the job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importexport_getStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(importexport_getStatusCmd).Standalone()

		importexport_getStatusCmd.Flags().String("apiversion", "", "")
		importexport_getStatusCmd.Flags().String("job-id", "", "")
		importexport_getStatusCmd.MarkFlagRequired("job-id")
	})
	importexportCmd.AddCommand(importexport_getStatusCmd)
}
