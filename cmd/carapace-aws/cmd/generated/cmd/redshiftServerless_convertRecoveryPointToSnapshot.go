package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_convertRecoveryPointToSnapshotCmd = &cobra.Command{
	Use:   "convert-recovery-point-to-snapshot",
	Short: "Converts a recovery point to a snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_convertRecoveryPointToSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftServerless_convertRecoveryPointToSnapshotCmd).Standalone()

		redshiftServerless_convertRecoveryPointToSnapshotCmd.Flags().String("recovery-point-id", "", "The unique identifier of the recovery point.")
		redshiftServerless_convertRecoveryPointToSnapshotCmd.Flags().String("retention-period", "", "How long to retain the snapshot.")
		redshiftServerless_convertRecoveryPointToSnapshotCmd.Flags().String("snapshot-name", "", "The name of the snapshot.")
		redshiftServerless_convertRecoveryPointToSnapshotCmd.Flags().String("tags", "", "An array of [Tag objects](https://docs.aws.amazon.com/redshift-serverless/latest/APIReference/API_Tag.html) to associate with the created snapshot.")
		redshiftServerless_convertRecoveryPointToSnapshotCmd.MarkFlagRequired("recovery-point-id")
		redshiftServerless_convertRecoveryPointToSnapshotCmd.MarkFlagRequired("snapshot-name")
	})
	redshiftServerlessCmd.AddCommand(redshiftServerless_convertRecoveryPointToSnapshotCmd)
}
