package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_listPodIdentityAssociationsCmd = &cobra.Command{
	Use:   "list-pod-identity-associations",
	Short: "List the EKS Pod Identity associations in a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_listPodIdentityAssociationsCmd).Standalone()

	eks_listPodIdentityAssociationsCmd.Flags().String("cluster-name", "", "The name of the cluster that the associations are in.")
	eks_listPodIdentityAssociationsCmd.Flags().String("max-results", "", "The maximum number of EKS Pod Identity association results returned by `ListPodIdentityAssociations` in paginated output.")
	eks_listPodIdentityAssociationsCmd.Flags().String("namespace", "", "The name of the Kubernetes namespace inside the cluster that the associations are in.")
	eks_listPodIdentityAssociationsCmd.Flags().String("next-token", "", "The `nextToken` value returned from a previous paginated `ListUpdates` request where `maxResults` was used and the results exceeded the value of that parameter.")
	eks_listPodIdentityAssociationsCmd.Flags().String("service-account", "", "The name of the Kubernetes service account that the associations use.")
	eks_listPodIdentityAssociationsCmd.MarkFlagRequired("cluster-name")
	eksCmd.AddCommand(eks_listPodIdentityAssociationsCmd)
}
