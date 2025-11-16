package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panorama_describePackageImportJobCmd = &cobra.Command{
	Use:   "describe-package-import-job",
	Short: "Returns information about a package import job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panorama_describePackageImportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(panorama_describePackageImportJobCmd).Standalone()

		panorama_describePackageImportJobCmd.Flags().String("job-id", "", "The job's ID.")
		panorama_describePackageImportJobCmd.MarkFlagRequired("job-id")
	})
	panoramaCmd.AddCommand(panorama_describePackageImportJobCmd)
}
