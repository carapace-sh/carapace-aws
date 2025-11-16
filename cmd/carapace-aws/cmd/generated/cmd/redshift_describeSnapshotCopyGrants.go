package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeSnapshotCopyGrantsCmd = &cobra.Command{
	Use:   "describe-snapshot-copy-grants",
	Short: "Returns a list of snapshot copy grants owned by the Amazon Web Services account in the destination region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeSnapshotCopyGrantsCmd).Standalone()

	redshift_describeSnapshotCopyGrantsCmd.Flags().String("marker", "", "An optional parameter that specifies the starting point to return a set of response records.")
	redshift_describeSnapshotCopyGrantsCmd.Flags().String("max-records", "", "The maximum number of response records to return in each call.")
	redshift_describeSnapshotCopyGrantsCmd.Flags().String("snapshot-copy-grant-name", "", "The name of the snapshot copy grant.")
	redshift_describeSnapshotCopyGrantsCmd.Flags().String("tag-keys", "", "A tag key or keys for which you want to return all matching resources that are associated with the specified key or keys.")
	redshift_describeSnapshotCopyGrantsCmd.Flags().String("tag-values", "", "A tag value or values for which you want to return all matching resources that are associated with the specified value or values.")
	redshiftCmd.AddCommand(redshift_describeSnapshotCopyGrantsCmd)
}
