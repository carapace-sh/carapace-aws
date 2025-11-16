package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_startImportJobCmd = &cobra.Command{
	Use:   "start-import-job",
	Short: "Start an asynchronous job to import Amazon Q in Connect resources from an uploaded source file.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_startImportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_startImportJobCmd).Standalone()

		qconnect_startImportJobCmd.Flags().String("client-token", "", "The tags used to organize, track, or control access for this resource.")
		qconnect_startImportJobCmd.Flags().String("external-source-configuration", "", "The configuration information of the external source that the resource data are imported from.")
		qconnect_startImportJobCmd.Flags().String("import-job-type", "", "The type of the import job.")
		qconnect_startImportJobCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
		qconnect_startImportJobCmd.Flags().String("metadata", "", "The metadata fields of the imported Amazon Q in Connect resources.")
		qconnect_startImportJobCmd.Flags().String("upload-id", "", "A pointer to the uploaded asset.")
		qconnect_startImportJobCmd.MarkFlagRequired("import-job-type")
		qconnect_startImportJobCmd.MarkFlagRequired("knowledge-base-id")
		qconnect_startImportJobCmd.MarkFlagRequired("upload-id")
	})
	qconnectCmd.AddCommand(qconnect_startImportJobCmd)
}
