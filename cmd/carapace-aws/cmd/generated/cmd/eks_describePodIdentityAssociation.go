package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_describePodIdentityAssociationCmd = &cobra.Command{
	Use:   "describe-pod-identity-association",
	Short: "Returns descriptive information about an EKS Pod Identity association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_describePodIdentityAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(eks_describePodIdentityAssociationCmd).Standalone()

		eks_describePodIdentityAssociationCmd.Flags().String("association-id", "", "The ID of the association that you want the description of.")
		eks_describePodIdentityAssociationCmd.Flags().String("cluster-name", "", "The name of the cluster that the association is in.")
		eks_describePodIdentityAssociationCmd.MarkFlagRequired("association-id")
		eks_describePodIdentityAssociationCmd.MarkFlagRequired("cluster-name")
	})
	eksCmd.AddCommand(eks_describePodIdentityAssociationCmd)
}
