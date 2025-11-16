package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efs_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from an EFS resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efs_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(efs_untagResourceCmd).Standalone()

		efs_untagResourceCmd.Flags().String("resource-id", "", "Specifies the EFS resource that you want to remove tags from.")
		efs_untagResourceCmd.Flags().String("tag-keys", "", "The keys of the key-value tag pairs that you want to remove from the specified EFS resource.")
		efs_untagResourceCmd.MarkFlagRequired("resource-id")
		efs_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	efsCmd.AddCommand(efs_untagResourceCmd)
}
