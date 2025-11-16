package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_createResourcePolicyStatementCmd = &cobra.Command{
	Use:   "create-resource-policy-statement",
	Short: "Adds a new resource policy statement to a bot or bot alias.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_createResourcePolicyStatementCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lexv2Models_createResourcePolicyStatementCmd).Standalone()

		lexv2Models_createResourcePolicyStatementCmd.Flags().String("action", "", "The Amazon Lex action that this policy either allows or denies.")
		lexv2Models_createResourcePolicyStatementCmd.Flags().String("condition", "", "Specifies a condition when the policy is in effect.")
		lexv2Models_createResourcePolicyStatementCmd.Flags().String("effect", "", "Determines whether the statement allows or denies access to the resource.")
		lexv2Models_createResourcePolicyStatementCmd.Flags().String("expected-revision-id", "", "The identifier of the revision of the policy to edit.")
		lexv2Models_createResourcePolicyStatementCmd.Flags().String("principal", "", "An IAM principal, such as an IAM user, IAM role, or Amazon Web Services services that is allowed or denied access to a resource.")
		lexv2Models_createResourcePolicyStatementCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the bot or bot alias that the resource policy is attached to.")
		lexv2Models_createResourcePolicyStatementCmd.Flags().String("statement-id", "", "The name of the statement.")
		lexv2Models_createResourcePolicyStatementCmd.MarkFlagRequired("action")
		lexv2Models_createResourcePolicyStatementCmd.MarkFlagRequired("effect")
		lexv2Models_createResourcePolicyStatementCmd.MarkFlagRequired("principal")
		lexv2Models_createResourcePolicyStatementCmd.MarkFlagRequired("resource-arn")
		lexv2Models_createResourcePolicyStatementCmd.MarkFlagRequired("statement-id")
	})
	lexv2ModelsCmd.AddCommand(lexv2Models_createResourcePolicyStatementCmd)
}
