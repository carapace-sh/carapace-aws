package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_addPermissionCmd = &cobra.Command{
	Use:   "add-permission",
	Short: "Grants a [principal](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_principal.html#Principal_specifying) permission to use a function.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_addPermissionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lambda_addPermissionCmd).Standalone()

		lambda_addPermissionCmd.Flags().String("action", "", "The action that the principal can use on the function.")
		lambda_addPermissionCmd.Flags().String("event-source-token", "", "For Alexa Smart Home functions, a token that the invoker must supply.")
		lambda_addPermissionCmd.Flags().String("function-name", "", "The name or ARN of the Lambda function, version, or alias.")
		lambda_addPermissionCmd.Flags().String("function-url-auth-type", "", "The type of authentication that your function URL uses.")
		lambda_addPermissionCmd.Flags().String("invoked-via-function-url", "", "Restricts the `lambda:InvokeFunction` action to function URL calls.")
		lambda_addPermissionCmd.Flags().String("principal", "", "The Amazon Web Services service, Amazon Web Services account, IAM user, or IAM role that invokes the function.")
		lambda_addPermissionCmd.Flags().String("principal-org-id", "", "The identifier for your organization in Organizations.")
		lambda_addPermissionCmd.Flags().String("qualifier", "", "Specify a version or alias to add permissions to a published version of the function.")
		lambda_addPermissionCmd.Flags().String("revision-id", "", "Update the policy only if the revision ID matches the ID that's specified.")
		lambda_addPermissionCmd.Flags().String("source-account", "", "For Amazon Web Services service, the ID of the Amazon Web Services account that owns the resource.")
		lambda_addPermissionCmd.Flags().String("source-arn", "", "For Amazon Web Services services, the ARN of the Amazon Web Services resource that invokes the function.")
		lambda_addPermissionCmd.Flags().String("statement-id", "", "A statement identifier that differentiates the statement from others in the same policy.")
		lambda_addPermissionCmd.MarkFlagRequired("action")
		lambda_addPermissionCmd.MarkFlagRequired("function-name")
		lambda_addPermissionCmd.MarkFlagRequired("principal")
		lambda_addPermissionCmd.MarkFlagRequired("statement-id")
	})
	lambdaCmd.AddCommand(lambda_addPermissionCmd)
}
