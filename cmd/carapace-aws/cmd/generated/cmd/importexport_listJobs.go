package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var importexport_listJobsCmd = &cobra.Command{
	Use:   "list-jobs",
	Short: "This operation returns the jobs associated with the requester.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importexport_listJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(importexport_listJobsCmd).Standalone()

		importexport_listJobsCmd.Flags().String("apiversion", "", "")
		importexport_listJobsCmd.Flags().String("marker", "", "")
		importexport_listJobsCmd.Flags().String("max-jobs", "", "")
	})
	importexportCmd.AddCommand(importexport_listJobsCmd)
}
