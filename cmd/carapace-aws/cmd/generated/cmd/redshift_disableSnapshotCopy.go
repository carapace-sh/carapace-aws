package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_disableSnapshotCopyCmd = &cobra.Command{
	Use:   "disable-snapshot-copy",
	Short: "Disables the automatic copying of snapshots from one region to another region for a specified cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_disableSnapshotCopyCmd).Standalone()

	redshift_disableSnapshotCopyCmd.Flags().String("cluster-identifier", "", "The unique identifier of the source cluster that you want to disable copying of snapshots to a destination region.")
	redshift_disableSnapshotCopyCmd.MarkFlagRequired("cluster-identifier")
	redshiftCmd.AddCommand(redshift_disableSnapshotCopyCmd)
}
