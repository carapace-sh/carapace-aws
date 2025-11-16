package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_listSnapshotCopyConfigurationsCmd = &cobra.Command{
	Use:   "list-snapshot-copy-configurations",
	Short: "Returns a list of snapshot copy configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_listSnapshotCopyConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftServerless_listSnapshotCopyConfigurationsCmd).Standalone()

		redshiftServerless_listSnapshotCopyConfigurationsCmd.Flags().String("max-results", "", "An optional parameter that specifies the maximum number of results to return.")
		redshiftServerless_listSnapshotCopyConfigurationsCmd.Flags().String("namespace-name", "", "The namespace from which to list all snapshot copy configurations.")
		redshiftServerless_listSnapshotCopyConfigurationsCmd.Flags().String("next-token", "", "If `nextToken` is returned, there are more results available.")
	})
	redshiftServerlessCmd.AddCommand(redshiftServerless_listSnapshotCopyConfigurationsCmd)
}
