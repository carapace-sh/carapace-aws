package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_getSamplingTargetsCmd = &cobra.Command{
	Use:   "get-sampling-targets",
	Short: "Requests a sampling quota for rules that the service is using to sample requests.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_getSamplingTargetsCmd).Standalone()

	xray_getSamplingTargetsCmd.Flags().String("sampling-boost-statistics-documents", "", "Information about rules that the service is using to boost sampling rate.")
	xray_getSamplingTargetsCmd.Flags().String("sampling-statistics-documents", "", "Information about rules that the service is using to sample requests.")
	xray_getSamplingTargetsCmd.MarkFlagRequired("sampling-statistics-documents")
	xrayCmd.AddCommand(xray_getSamplingTargetsCmd)
}
