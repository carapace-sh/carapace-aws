package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_createSbomExportCmd = &cobra.Command{
	Use:   "create-sbom-export",
	Short: "Creates a software bill of materials (SBOM) report.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_createSbomExportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_createSbomExportCmd).Standalone()

		inspector2_createSbomExportCmd.Flags().String("report-format", "", "The output format for the software bill of materials (SBOM) report.")
		inspector2_createSbomExportCmd.Flags().String("resource-filter-criteria", "", "The resource filter criteria for the software bill of materials (SBOM) report.")
		inspector2_createSbomExportCmd.Flags().String("s3-destination", "", "Contains details of the Amazon S3 bucket and KMS key used to export findings.")
		inspector2_createSbomExportCmd.MarkFlagRequired("report-format")
		inspector2_createSbomExportCmd.MarkFlagRequired("s3-destination")
	})
	inspector2Cmd.AddCommand(inspector2_createSbomExportCmd)
}
