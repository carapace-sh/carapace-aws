package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_listImportsCmd = &cobra.Command{
	Use:   "list-imports",
	Short: "Returns information on all imports, or a select set of imports by `ImportStatus` or `Destination`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_listImportsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudtrail_listImportsCmd).Standalone()

		cloudtrail_listImportsCmd.Flags().String("destination", "", "The ARN of the destination event data store.")
		cloudtrail_listImportsCmd.Flags().String("import-status", "", "The status of the import.")
		cloudtrail_listImportsCmd.Flags().String("max-results", "", "The maximum number of imports to display on a single page.")
		cloudtrail_listImportsCmd.Flags().String("next-token", "", "A token you can use to get the next page of import results.")
	})
	cloudtrailCmd.AddCommand(cloudtrail_listImportsCmd)
}
