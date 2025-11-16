package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ram_promoteResourceShareCreatedFromPolicyCmd = &cobra.Command{
	Use:   "promote-resource-share-created-from-policy",
	Short: "When you attach a resource-based policy to a resource, RAM automatically creates a resource share of `featureSet`=`CREATED_FROM_POLICY` with a managed permission that has the same IAM permissions as the original resource-based policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ram_promoteResourceShareCreatedFromPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ram_promoteResourceShareCreatedFromPolicyCmd).Standalone()

		ram_promoteResourceShareCreatedFromPolicyCmd.Flags().String("resource-share-arn", "", "Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the resource share to promote.")
		ram_promoteResourceShareCreatedFromPolicyCmd.MarkFlagRequired("resource-share-arn")
	})
	ramCmd.AddCommand(ram_promoteResourceShareCreatedFromPolicyCmd)
}
