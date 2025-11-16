package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_describeImageReplicationStatusCmd = &cobra.Command{
	Use:   "describe-image-replication-status",
	Short: "Returns the replication status for a specified image.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_describeImageReplicationStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecr_describeImageReplicationStatusCmd).Standalone()

		ecr_describeImageReplicationStatusCmd.Flags().String("image-id", "", "")
		ecr_describeImageReplicationStatusCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID associated with the registry.")
		ecr_describeImageReplicationStatusCmd.Flags().String("repository-name", "", "The name of the repository that the image is in.")
		ecr_describeImageReplicationStatusCmd.MarkFlagRequired("image-id")
		ecr_describeImageReplicationStatusCmd.MarkFlagRequired("repository-name")
	})
	ecrCmd.AddCommand(ecr_describeImageReplicationStatusCmd)
}
