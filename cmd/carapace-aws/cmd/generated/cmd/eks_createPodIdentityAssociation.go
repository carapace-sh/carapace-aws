package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_createPodIdentityAssociationCmd = &cobra.Command{
	Use:   "create-pod-identity-association",
	Short: "Creates an EKS Pod Identity association between a service account in an Amazon EKS cluster and an IAM role with *EKS Pod Identity*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_createPodIdentityAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(eks_createPodIdentityAssociationCmd).Standalone()

		eks_createPodIdentityAssociationCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		eks_createPodIdentityAssociationCmd.Flags().String("cluster-name", "", "The name of the cluster to create the EKS Pod Identity association in.")
		eks_createPodIdentityAssociationCmd.Flags().String("disable-session-tags", "", "Disable the automatic sessions tags that are appended by EKS Pod Identity.")
		eks_createPodIdentityAssociationCmd.Flags().String("namespace", "", "The name of the Kubernetes namespace inside the cluster to create the EKS Pod Identity association in.")
		eks_createPodIdentityAssociationCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the IAM role to associate with the service account.")
		eks_createPodIdentityAssociationCmd.Flags().String("service-account", "", "The name of the Kubernetes service account inside the cluster to associate the IAM credentials with.")
		eks_createPodIdentityAssociationCmd.Flags().String("tags", "", "Metadata that assists with categorization and organization.")
		eks_createPodIdentityAssociationCmd.Flags().String("target-role-arn", "", "The Amazon Resource Name (ARN) of the target IAM role to associate with the service account.")
		eks_createPodIdentityAssociationCmd.MarkFlagRequired("cluster-name")
		eks_createPodIdentityAssociationCmd.MarkFlagRequired("namespace")
		eks_createPodIdentityAssociationCmd.MarkFlagRequired("role-arn")
		eks_createPodIdentityAssociationCmd.MarkFlagRequired("service-account")
	})
	eksCmd.AddCommand(eks_createPodIdentityAssociationCmd)
}
