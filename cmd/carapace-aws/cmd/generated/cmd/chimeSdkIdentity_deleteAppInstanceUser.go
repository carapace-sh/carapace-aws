package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkIdentity_deleteAppInstanceUserCmd = &cobra.Command{
	Use:   "delete-app-instance-user",
	Short: "Deletes an `AppInstanceUser`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkIdentity_deleteAppInstanceUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkIdentity_deleteAppInstanceUserCmd).Standalone()

		chimeSdkIdentity_deleteAppInstanceUserCmd.Flags().String("app-instance-user-arn", "", "The ARN of the user request being deleted.")
		chimeSdkIdentity_deleteAppInstanceUserCmd.MarkFlagRequired("app-instance-user-arn")
	})
	chimeSdkIdentityCmd.AddCommand(chimeSdkIdentity_deleteAppInstanceUserCmd)
}
