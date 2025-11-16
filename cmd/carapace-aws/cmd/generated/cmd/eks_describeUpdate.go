package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_describeUpdateCmd = &cobra.Command{
	Use:   "describe-update",
	Short: "Describes an update to an Amazon EKS resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_describeUpdateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(eks_describeUpdateCmd).Standalone()

		eks_describeUpdateCmd.Flags().String("addon-name", "", "The name of the add-on.")
		eks_describeUpdateCmd.Flags().String("name", "", "The name of the Amazon EKS cluster associated with the update.")
		eks_describeUpdateCmd.Flags().String("nodegroup-name", "", "The name of the Amazon EKS node group associated with the update.")
		eks_describeUpdateCmd.Flags().String("update-id", "", "The ID of the update to describe.")
		eks_describeUpdateCmd.MarkFlagRequired("name")
		eks_describeUpdateCmd.MarkFlagRequired("update-id")
	})
	eksCmd.AddCommand(eks_describeUpdateCmd)
}
