package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_deletePodIdentityAssociationCmd = &cobra.Command{
	Use:   "delete-pod-identity-association",
	Short: "Deletes a EKS Pod Identity association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_deletePodIdentityAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(eks_deletePodIdentityAssociationCmd).Standalone()

		eks_deletePodIdentityAssociationCmd.Flags().String("association-id", "", "The ID of the association to be deleted.")
		eks_deletePodIdentityAssociationCmd.Flags().String("cluster-name", "", "The cluster name that")
		eks_deletePodIdentityAssociationCmd.MarkFlagRequired("association-id")
		eks_deletePodIdentityAssociationCmd.MarkFlagRequired("cluster-name")
	})
	eksCmd.AddCommand(eks_deletePodIdentityAssociationCmd)
}
