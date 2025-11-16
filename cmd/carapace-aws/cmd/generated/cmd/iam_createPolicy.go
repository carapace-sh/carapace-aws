package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_createPolicyCmd = &cobra.Command{
	Use:   "create-policy",
	Short: "Creates a new managed policy for your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_createPolicyCmd).Standalone()

	iam_createPolicyCmd.Flags().String("description", "", "A friendly description of the policy.")
	iam_createPolicyCmd.Flags().String("path", "", "The path for the policy.")
	iam_createPolicyCmd.Flags().String("policy-document", "", "The JSON policy document that you want to use as the content for the new policy.")
	iam_createPolicyCmd.Flags().String("policy-name", "", "The friendly name of the policy.")
	iam_createPolicyCmd.Flags().String("tags", "", "A list of tags that you want to attach to the new IAM customer managed policy.")
	iam_createPolicyCmd.MarkFlagRequired("policy-document")
	iam_createPolicyCmd.MarkFlagRequired("policy-name")
	iamCmd.AddCommand(iam_createPolicyCmd)
}
