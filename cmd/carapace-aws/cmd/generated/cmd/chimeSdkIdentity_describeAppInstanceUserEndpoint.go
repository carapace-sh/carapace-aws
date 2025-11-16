package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkIdentity_describeAppInstanceUserEndpointCmd = &cobra.Command{
	Use:   "describe-app-instance-user-endpoint",
	Short: "Returns the full details of an `AppInstanceUserEndpoint`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkIdentity_describeAppInstanceUserEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkIdentity_describeAppInstanceUserEndpointCmd).Standalone()

		chimeSdkIdentity_describeAppInstanceUserEndpointCmd.Flags().String("app-instance-user-arn", "", "The ARN of the `AppInstanceUser`.")
		chimeSdkIdentity_describeAppInstanceUserEndpointCmd.Flags().String("endpoint-id", "", "The unique identifier of the `AppInstanceUserEndpoint`.")
		chimeSdkIdentity_describeAppInstanceUserEndpointCmd.MarkFlagRequired("app-instance-user-arn")
		chimeSdkIdentity_describeAppInstanceUserEndpointCmd.MarkFlagRequired("endpoint-id")
	})
	chimeSdkIdentityCmd.AddCommand(chimeSdkIdentity_describeAppInstanceUserEndpointCmd)
}
