package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_listSnapshotsCmd = &cobra.Command{
	Use:   "list-snapshots",
	Short: "Returns a list of snapshots.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_listSnapshotsCmd).Standalone()

	redshiftServerless_listSnapshotsCmd.Flags().String("end-time", "", "The timestamp showing when the snapshot creation finished.")
	redshiftServerless_listSnapshotsCmd.Flags().String("max-results", "", "An optional parameter that specifies the maximum number of results to return.")
	redshiftServerless_listSnapshotsCmd.Flags().String("namespace-arn", "", "The Amazon Resource Name (ARN) of the namespace from which to list all snapshots.")
	redshiftServerless_listSnapshotsCmd.Flags().String("namespace-name", "", "The namespace from which to list all snapshots.")
	redshiftServerless_listSnapshotsCmd.Flags().String("next-token", "", "If `nextToken` is returned, there are more results available.")
	redshiftServerless_listSnapshotsCmd.Flags().String("owner-account", "", "The owner Amazon Web Services account of the snapshot.")
	redshiftServerless_listSnapshotsCmd.Flags().String("start-time", "", "The time when the creation of the snapshot was initiated.")
	redshiftServerlessCmd.AddCommand(redshiftServerless_listSnapshotsCmd)
}
