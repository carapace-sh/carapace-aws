package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_deleteAddonCmd = &cobra.Command{
	Use:   "delete-addon",
	Short: "Deletes an Amazon EKS add-on.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_deleteAddonCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(eks_deleteAddonCmd).Standalone()

		eks_deleteAddonCmd.Flags().String("addon-name", "", "The name of the add-on.")
		eks_deleteAddonCmd.Flags().String("cluster-name", "", "The name of your cluster.")
		eks_deleteAddonCmd.Flags().Bool("no-preserve", false, "Specifying this option preserves the add-on software on your cluster but Amazon EKS stops managing any settings for the add-on.")
		eks_deleteAddonCmd.Flags().Bool("preserve", false, "Specifying this option preserves the add-on software on your cluster but Amazon EKS stops managing any settings for the add-on.")
		eks_deleteAddonCmd.MarkFlagRequired("addon-name")
		eks_deleteAddonCmd.MarkFlagRequired("cluster-name")
		eks_deleteAddonCmd.Flag("no-preserve").Hidden = true
	})
	eksCmd.AddCommand(eks_deleteAddonCmd)
}
