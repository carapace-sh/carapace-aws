package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_listImportFailuresCmd = &cobra.Command{
	Use:   "list-import-failures",
	Short: "Returns a list of failures for the specified import.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_listImportFailuresCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudtrail_listImportFailuresCmd).Standalone()

		cloudtrail_listImportFailuresCmd.Flags().String("import-id", "", "The ID of the import.")
		cloudtrail_listImportFailuresCmd.Flags().String("max-results", "", "The maximum number of failures to display on a single page.")
		cloudtrail_listImportFailuresCmd.Flags().String("next-token", "", "A token you can use to get the next page of import failures.")
		cloudtrail_listImportFailuresCmd.MarkFlagRequired("import-id")
	})
	cloudtrailCmd.AddCommand(cloudtrail_listImportFailuresCmd)
}
