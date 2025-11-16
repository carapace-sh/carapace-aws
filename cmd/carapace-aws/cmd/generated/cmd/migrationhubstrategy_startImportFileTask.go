package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhubstrategy_startImportFileTaskCmd = &cobra.Command{
	Use:   "start-import-file-task",
	Short: "Starts a file import.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhubstrategy_startImportFileTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(migrationhubstrategy_startImportFileTaskCmd).Standalone()

		migrationhubstrategy_startImportFileTaskCmd.Flags().String("data-source-type", "", "Specifies the source that the servers are coming from.")
		migrationhubstrategy_startImportFileTaskCmd.Flags().String("group-id", "", "Groups the resources in the import file together with a unique name.")
		migrationhubstrategy_startImportFileTaskCmd.Flags().String("name", "", "A descriptive name for the request.")
		migrationhubstrategy_startImportFileTaskCmd.Flags().String("s3-bucket", "", "The S3 bucket where the import file is located.")
		migrationhubstrategy_startImportFileTaskCmd.Flags().String("s3bucket-for-report-data", "", "The S3 bucket where Strategy Recommendations uploads import results.")
		migrationhubstrategy_startImportFileTaskCmd.Flags().String("s3key", "", "The Amazon S3 key name of the import file.")
		migrationhubstrategy_startImportFileTaskCmd.MarkFlagRequired("name")
		migrationhubstrategy_startImportFileTaskCmd.MarkFlagRequired("s3-bucket")
		migrationhubstrategy_startImportFileTaskCmd.MarkFlagRequired("s3key")
	})
	migrationhubstrategyCmd.AddCommand(migrationhubstrategy_startImportFileTaskCmd)
}
