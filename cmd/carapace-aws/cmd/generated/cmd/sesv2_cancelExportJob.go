package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_cancelExportJobCmd = &cobra.Command{
	Use:   "cancel-export-job",
	Short: "Cancels an export job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_cancelExportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_cancelExportJobCmd).Standalone()

		sesv2_cancelExportJobCmd.Flags().String("job-id", "", "The export job ID.")
		sesv2_cancelExportJobCmd.MarkFlagRequired("job-id")
	})
	sesv2Cmd.AddCommand(sesv2_cancelExportJobCmd)
}
