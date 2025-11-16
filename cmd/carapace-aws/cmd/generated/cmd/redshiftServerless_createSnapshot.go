package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_createSnapshotCmd = &cobra.Command{
	Use:   "create-snapshot",
	Short: "Creates a snapshot of all databases in a namespace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_createSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftServerless_createSnapshotCmd).Standalone()

		redshiftServerless_createSnapshotCmd.Flags().String("namespace-name", "", "The namespace to create a snapshot for.")
		redshiftServerless_createSnapshotCmd.Flags().String("retention-period", "", "How long to retain the created snapshot.")
		redshiftServerless_createSnapshotCmd.Flags().String("snapshot-name", "", "The name of the snapshot.")
		redshiftServerless_createSnapshotCmd.Flags().String("tags", "", "An array of [Tag objects](https://docs.aws.amazon.com/redshift-serverless/latest/APIReference/API_Tag.html) to associate with the snapshot.")
		redshiftServerless_createSnapshotCmd.MarkFlagRequired("namespace-name")
		redshiftServerless_createSnapshotCmd.MarkFlagRequired("snapshot-name")
	})
	redshiftServerlessCmd.AddCommand(redshiftServerless_createSnapshotCmd)
}
