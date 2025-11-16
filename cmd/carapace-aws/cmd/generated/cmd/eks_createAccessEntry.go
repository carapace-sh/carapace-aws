package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_createAccessEntryCmd = &cobra.Command{
	Use:   "create-access-entry",
	Short: "Creates an access entry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_createAccessEntryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(eks_createAccessEntryCmd).Standalone()

		eks_createAccessEntryCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		eks_createAccessEntryCmd.Flags().String("cluster-name", "", "The name of your cluster.")
		eks_createAccessEntryCmd.Flags().String("kubernetes-groups", "", "The value for `name` that you've specified for `kind: Group` as a `subject` in a Kubernetes `RoleBinding` or `ClusterRoleBinding` object.")
		eks_createAccessEntryCmd.Flags().String("principal-arn", "", "The ARN of the IAM principal for the `AccessEntry`.")
		eks_createAccessEntryCmd.Flags().String("tags", "", "Metadata that assists with categorization and organization.")
		eks_createAccessEntryCmd.Flags().String("type", "", "The type of the new access entry.")
		eks_createAccessEntryCmd.Flags().String("username", "", "The username to authenticate to Kubernetes with.")
		eks_createAccessEntryCmd.MarkFlagRequired("cluster-name")
		eks_createAccessEntryCmd.MarkFlagRequired("principal-arn")
	})
	eksCmd.AddCommand(eks_createAccessEntryCmd)
}
