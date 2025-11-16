package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_getSbomExportCmd = &cobra.Command{
	Use:   "get-sbom-export",
	Short: "Gets details of a software bill of materials (SBOM) report.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_getSbomExportCmd).Standalone()

	inspector2_getSbomExportCmd.Flags().String("report-id", "", "The report ID of the SBOM export to get details for.")
	inspector2_getSbomExportCmd.MarkFlagRequired("report-id")
	inspector2Cmd.AddCommand(inspector2_getSbomExportCmd)
}
