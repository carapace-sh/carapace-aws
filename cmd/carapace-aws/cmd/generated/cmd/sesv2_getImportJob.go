package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_getImportJobCmd = &cobra.Command{
	Use:   "get-import-job",
	Short: "Provides information about an import job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_getImportJobCmd).Standalone()

	sesv2_getImportJobCmd.Flags().String("job-id", "", "The ID of the import job.")
	sesv2_getImportJobCmd.MarkFlagRequired("job-id")
	sesv2Cmd.AddCommand(sesv2_getImportJobCmd)
}
