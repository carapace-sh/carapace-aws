package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkIdentity_describeAppInstanceAdminCmd = &cobra.Command{
	Use:   "describe-app-instance-admin",
	Short: "Returns the full details of an `AppInstanceAdmin`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkIdentity_describeAppInstanceAdminCmd).Standalone()

	chimeSdkIdentity_describeAppInstanceAdminCmd.Flags().String("app-instance-admin-arn", "", "The ARN of the `AppInstanceAdmin`.")
	chimeSdkIdentity_describeAppInstanceAdminCmd.Flags().String("app-instance-arn", "", "The ARN of the `AppInstance`.")
	chimeSdkIdentity_describeAppInstanceAdminCmd.MarkFlagRequired("app-instance-admin-arn")
	chimeSdkIdentity_describeAppInstanceAdminCmd.MarkFlagRequired("app-instance-arn")
	chimeSdkIdentityCmd.AddCommand(chimeSdkIdentity_describeAppInstanceAdminCmd)
}
