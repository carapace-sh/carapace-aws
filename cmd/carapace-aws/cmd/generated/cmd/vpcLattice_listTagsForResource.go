package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags for the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_listTagsForResourceCmd).Standalone()

		vpcLattice_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		vpcLattice_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_listTagsForResourceCmd)
}
