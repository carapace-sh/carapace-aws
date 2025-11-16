package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdbElastic_copyClusterSnapshotCmd = &cobra.Command{
	Use:   "copy-cluster-snapshot",
	Short: "Copies a snapshot of an elastic cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdbElastic_copyClusterSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdbElastic_copyClusterSnapshotCmd).Standalone()

		docdbElastic_copyClusterSnapshotCmd.Flags().Bool("copy-tags", false, "Set to `true` to copy all tags from the source cluster snapshot to the target elastic cluster snapshot.")
		docdbElastic_copyClusterSnapshotCmd.Flags().String("kms-key-id", "", "The Amazon Web Services KMS key ID for an encrypted elastic cluster snapshot.")
		docdbElastic_copyClusterSnapshotCmd.Flags().Bool("no-copy-tags", false, "Set to `true` to copy all tags from the source cluster snapshot to the target elastic cluster snapshot.")
		docdbElastic_copyClusterSnapshotCmd.Flags().String("snapshot-arn", "", "The Amazon Resource Name (ARN) identifier of the elastic cluster snapshot.")
		docdbElastic_copyClusterSnapshotCmd.Flags().String("tags", "", "The tags to be assigned to the elastic cluster snapshot.")
		docdbElastic_copyClusterSnapshotCmd.Flags().String("target-snapshot-name", "", "The identifier of the new elastic cluster snapshot to create from the source cluster snapshot.")
		docdbElastic_copyClusterSnapshotCmd.Flag("no-copy-tags").Hidden = true
		docdbElastic_copyClusterSnapshotCmd.MarkFlagRequired("snapshot-arn")
		docdbElastic_copyClusterSnapshotCmd.MarkFlagRequired("target-snapshot-name")
	})
	docdbElasticCmd.AddCommand(docdbElastic_copyClusterSnapshotCmd)
}
