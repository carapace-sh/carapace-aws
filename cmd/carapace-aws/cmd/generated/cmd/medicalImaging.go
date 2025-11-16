package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medicalImagingCmd = &cobra.Command{
	Use:   "medical-imaging",
	Short: "This is the *AWS HealthImaging API Reference*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medicalImagingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medicalImagingCmd).Standalone()

	})
	rootCmd.AddCommand(medicalImagingCmd)
}
