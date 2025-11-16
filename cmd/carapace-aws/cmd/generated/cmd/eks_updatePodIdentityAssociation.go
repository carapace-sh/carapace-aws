package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_updatePodIdentityAssociationCmd = &cobra.Command{
	Use:   "update-pod-identity-association",
	Short: "Updates a EKS Pod Identity association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_updatePodIdentityAssociationCmd).Standalone()

	eks_updatePodIdentityAssociationCmd.Flags().String("association-id", "", "The ID of the association to be updated.")
	eks_updatePodIdentityAssociationCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	eks_updatePodIdentityAssociationCmd.Flags().String("cluster-name", "", "The name of the cluster that you want to update the association in.")
	eks_updatePodIdentityAssociationCmd.Flags().String("disable-session-tags", "", "Disable the automatic sessions tags that are appended by EKS Pod Identity.")
	eks_updatePodIdentityAssociationCmd.Flags().String("role-arn", "", "The new IAM role to change in the association.")
	eks_updatePodIdentityAssociationCmd.Flags().String("target-role-arn", "", "The Amazon Resource Name (ARN) of the target IAM role to associate with the service account.")
	eks_updatePodIdentityAssociationCmd.MarkFlagRequired("association-id")
	eks_updatePodIdentityAssociationCmd.MarkFlagRequired("cluster-name")
	eksCmd.AddCommand(eks_updatePodIdentityAssociationCmd)
}
