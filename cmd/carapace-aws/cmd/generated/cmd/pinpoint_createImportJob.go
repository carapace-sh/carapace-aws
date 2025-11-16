package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_createImportJobCmd = &cobra.Command{
	Use:   "create-import-job",
	Short: "Creates an import job for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_createImportJobCmd).Standalone()

	pinpoint_createImportJobCmd.Flags().String("application-id", "", "The unique identifier for the application.")
	pinpoint_createImportJobCmd.Flags().String("import-job-request", "", "")
	pinpoint_createImportJobCmd.MarkFlagRequired("application-id")
	pinpoint_createImportJobCmd.MarkFlagRequired("import-job-request")
	pinpointCmd.AddCommand(pinpoint_createImportJobCmd)
}
