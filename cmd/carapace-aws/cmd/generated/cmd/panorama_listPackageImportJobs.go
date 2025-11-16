package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panorama_listPackageImportJobsCmd = &cobra.Command{
	Use:   "list-package-import-jobs",
	Short: "Returns a list of package import jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panorama_listPackageImportJobsCmd).Standalone()

	panorama_listPackageImportJobsCmd.Flags().String("max-results", "", "The maximum number of package import jobs to return in one page of results.")
	panorama_listPackageImportJobsCmd.Flags().String("next-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
	panoramaCmd.AddCommand(panorama_listPackageImportJobsCmd)
}
