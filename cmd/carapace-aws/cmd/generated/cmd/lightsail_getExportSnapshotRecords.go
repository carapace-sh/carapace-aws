package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getExportSnapshotRecordsCmd = &cobra.Command{
	Use:   "get-export-snapshot-records",
	Short: "Returns all export snapshot records created as a result of the `export snapshot` operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getExportSnapshotRecordsCmd).Standalone()

	lightsail_getExportSnapshotRecordsCmd.Flags().String("page-token", "", "The token to advance to the next page of results from your request.")
	lightsailCmd.AddCommand(lightsail_getExportSnapshotRecordsCmd)
}
