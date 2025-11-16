package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ram_promotePermissionCreatedFromPolicyCmd = &cobra.Command{
	Use:   "promote-permission-created-from-policy",
	Short: "When you attach a resource-based policy to a resource, RAM automatically creates a resource share of `featureSet`=`CREATED_FROM_POLICY` with a managed permission that has the same IAM permissions as the original resource-based policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ram_promotePermissionCreatedFromPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ram_promotePermissionCreatedFromPolicyCmd).Standalone()

		ram_promotePermissionCreatedFromPolicyCmd.Flags().String("client-token", "", "Specifies a unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ram_promotePermissionCreatedFromPolicyCmd.Flags().String("name", "", "Specifies a name for the promoted customer managed permission.")
		ram_promotePermissionCreatedFromPolicyCmd.Flags().String("permission-arn", "", "Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the `CREATED_FROM_POLICY` permission that you want to promote.")
		ram_promotePermissionCreatedFromPolicyCmd.MarkFlagRequired("name")
		ram_promotePermissionCreatedFromPolicyCmd.MarkFlagRequired("permission-arn")
	})
	ramCmd.AddCommand(ram_promotePermissionCreatedFromPolicyCmd)
}
