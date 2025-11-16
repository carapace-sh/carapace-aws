package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_startImportJobCmd = &cobra.Command{
	Use:   "start-import-job",
	Short: "Start an asynchronous job to import Wisdom resources from an uploaded source file.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_startImportJobCmd).Standalone()

	wisdom_startImportJobCmd.Flags().String("client-token", "", "The tags used to organize, track, or control access for this resource.")
	wisdom_startImportJobCmd.Flags().String("external-source-configuration", "", "The configuration information of the external source that the resource data are imported from.")
	wisdom_startImportJobCmd.Flags().String("import-job-type", "", "The type of the import job.")
	wisdom_startImportJobCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
	wisdom_startImportJobCmd.Flags().String("metadata", "", "The metadata fields of the imported Wisdom resources.")
	wisdom_startImportJobCmd.Flags().String("upload-id", "", "A pointer to the uploaded asset.")
	wisdom_startImportJobCmd.MarkFlagRequired("import-job-type")
	wisdom_startImportJobCmd.MarkFlagRequired("knowledge-base-id")
	wisdom_startImportJobCmd.MarkFlagRequired("upload-id")
	wisdomCmd.AddCommand(wisdom_startImportJobCmd)
}
