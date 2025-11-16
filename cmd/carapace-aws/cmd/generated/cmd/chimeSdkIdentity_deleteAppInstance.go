package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkIdentity_deleteAppInstanceCmd = &cobra.Command{
	Use:   "delete-app-instance",
	Short: "Deletes an `AppInstance` and all associated data asynchronously.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkIdentity_deleteAppInstanceCmd).Standalone()

	chimeSdkIdentity_deleteAppInstanceCmd.Flags().String("app-instance-arn", "", "The ARN of the `AppInstance`.")
	chimeSdkIdentity_deleteAppInstanceCmd.MarkFlagRequired("app-instance-arn")
	chimeSdkIdentityCmd.AddCommand(chimeSdkIdentity_deleteAppInstanceCmd)
}
