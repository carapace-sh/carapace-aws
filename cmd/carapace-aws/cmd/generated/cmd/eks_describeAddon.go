package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_describeAddonCmd = &cobra.Command{
	Use:   "describe-addon",
	Short: "Describes an Amazon EKS add-on.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_describeAddonCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(eks_describeAddonCmd).Standalone()

		eks_describeAddonCmd.Flags().String("addon-name", "", "The name of the add-on.")
		eks_describeAddonCmd.Flags().String("cluster-name", "", "The name of your cluster.")
		eks_describeAddonCmd.MarkFlagRequired("addon-name")
		eks_describeAddonCmd.MarkFlagRequired("cluster-name")
	})
	eksCmd.AddCommand(eks_describeAddonCmd)
}
