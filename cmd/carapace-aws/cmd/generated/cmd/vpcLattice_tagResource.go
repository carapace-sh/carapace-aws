package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds the specified tags to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_tagResourceCmd).Standalone()

		vpcLattice_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		vpcLattice_tagResourceCmd.Flags().String("tags", "", "The tags for the resource.")
		vpcLattice_tagResourceCmd.MarkFlagRequired("resource-arn")
		vpcLattice_tagResourceCmd.MarkFlagRequired("tags")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_tagResourceCmd)
}
