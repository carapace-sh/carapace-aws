package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_updateAccessEntryCmd = &cobra.Command{
	Use:   "update-access-entry",
	Short: "Updates an access entry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_updateAccessEntryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(eks_updateAccessEntryCmd).Standalone()

		eks_updateAccessEntryCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		eks_updateAccessEntryCmd.Flags().String("cluster-name", "", "The name of your cluster.")
		eks_updateAccessEntryCmd.Flags().String("kubernetes-groups", "", "The value for `name` that you've specified for `kind: Group` as a `subject` in a Kubernetes `RoleBinding` or `ClusterRoleBinding` object.")
		eks_updateAccessEntryCmd.Flags().String("principal-arn", "", "The ARN of the IAM principal for the `AccessEntry`.")
		eks_updateAccessEntryCmd.Flags().String("username", "", "The username to authenticate to Kubernetes with.")
		eks_updateAccessEntryCmd.MarkFlagRequired("cluster-name")
		eks_updateAccessEntryCmd.MarkFlagRequired("principal-arn")
	})
	eksCmd.AddCommand(eks_updateAccessEntryCmd)
}
