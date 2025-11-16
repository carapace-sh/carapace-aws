package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panorama_createPackageImportJobCmd = &cobra.Command{
	Use:   "create-package-import-job",
	Short: "Imports a node package.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panorama_createPackageImportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(panorama_createPackageImportJobCmd).Standalone()

		panorama_createPackageImportJobCmd.Flags().String("client-token", "", "A client token for the package import job.")
		panorama_createPackageImportJobCmd.Flags().String("input-config", "", "An input config for the package import job.")
		panorama_createPackageImportJobCmd.Flags().String("job-tags", "", "Tags for the package import job.")
		panorama_createPackageImportJobCmd.Flags().String("job-type", "", "A job type for the package import job.")
		panorama_createPackageImportJobCmd.Flags().String("output-config", "", "An output config for the package import job.")
		panorama_createPackageImportJobCmd.MarkFlagRequired("client-token")
		panorama_createPackageImportJobCmd.MarkFlagRequired("input-config")
		panorama_createPackageImportJobCmd.MarkFlagRequired("job-type")
		panorama_createPackageImportJobCmd.MarkFlagRequired("output-config")
	})
	panoramaCmd.AddCommand(panorama_createPackageImportJobCmd)
}
