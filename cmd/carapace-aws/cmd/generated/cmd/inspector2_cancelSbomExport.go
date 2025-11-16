package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_cancelSbomExportCmd = &cobra.Command{
	Use:   "cancel-sbom-export",
	Short: "Cancels a software bill of materials (SBOM) report.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_cancelSbomExportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_cancelSbomExportCmd).Standalone()

		inspector2_cancelSbomExportCmd.Flags().String("report-id", "", "The report ID of the SBOM export to cancel.")
		inspector2_cancelSbomExportCmd.MarkFlagRequired("report-id")
	})
	inspector2Cmd.AddCommand(inspector2_cancelSbomExportCmd)
}
