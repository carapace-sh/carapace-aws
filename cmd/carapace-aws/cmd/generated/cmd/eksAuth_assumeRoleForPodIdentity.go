package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eksAuth_assumeRoleForPodIdentityCmd = &cobra.Command{
	Use:   "assume-role-for-pod-identity",
	Short: "The Amazon EKS Auth API and the `AssumeRoleForPodIdentity` action are only used by the EKS Pod Identity Agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eksAuth_assumeRoleForPodIdentityCmd).Standalone()

	eksAuth_assumeRoleForPodIdentityCmd.Flags().String("cluster-name", "", "The name of the cluster for the request.")
	eksAuth_assumeRoleForPodIdentityCmd.Flags().String("token", "", "The token of the Kubernetes service account for the pod.")
	eksAuth_assumeRoleForPodIdentityCmd.MarkFlagRequired("cluster-name")
	eksAuth_assumeRoleForPodIdentityCmd.MarkFlagRequired("token")
	eksAuthCmd.AddCommand(eksAuth_assumeRoleForPodIdentityCmd)
}
