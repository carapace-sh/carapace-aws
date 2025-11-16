package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkIdentity_updateAppInstanceUserCmd = &cobra.Command{
	Use:   "update-app-instance-user",
	Short: "Updates the details of an `AppInstanceUser`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkIdentity_updateAppInstanceUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkIdentity_updateAppInstanceUserCmd).Standalone()

		chimeSdkIdentity_updateAppInstanceUserCmd.Flags().String("app-instance-user-arn", "", "The ARN of the `AppInstanceUser`.")
		chimeSdkIdentity_updateAppInstanceUserCmd.Flags().String("metadata", "", "The metadata of the `AppInstanceUser`.")
		chimeSdkIdentity_updateAppInstanceUserCmd.Flags().String("name", "", "The name of the `AppInstanceUser`.")
		chimeSdkIdentity_updateAppInstanceUserCmd.MarkFlagRequired("app-instance-user-arn")
		chimeSdkIdentity_updateAppInstanceUserCmd.MarkFlagRequired("metadata")
		chimeSdkIdentity_updateAppInstanceUserCmd.MarkFlagRequired("name")
	})
	chimeSdkIdentityCmd.AddCommand(chimeSdkIdentity_updateAppInstanceUserCmd)
}
