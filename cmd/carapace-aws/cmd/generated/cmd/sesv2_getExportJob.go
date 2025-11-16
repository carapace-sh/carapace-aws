package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_getExportJobCmd = &cobra.Command{
	Use:   "get-export-job",
	Short: "Provides information about an export job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_getExportJobCmd).Standalone()

	sesv2_getExportJobCmd.Flags().String("job-id", "", "The export job ID.")
	sesv2_getExportJobCmd.MarkFlagRequired("job-id")
	sesv2Cmd.AddCommand(sesv2_getExportJobCmd)
}
