package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the specified tags from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_untagResourceCmd).Standalone()

		vpcLattice_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		vpcLattice_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys of the tags to remove.")
		vpcLattice_untagResourceCmd.MarkFlagRequired("resource-arn")
		vpcLattice_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_untagResourceCmd)
}
