package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkIdentity_listAppInstanceAdminsCmd = &cobra.Command{
	Use:   "list-app-instance-admins",
	Short: "Returns a list of the administrators in the `AppInstance`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkIdentity_listAppInstanceAdminsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkIdentity_listAppInstanceAdminsCmd).Standalone()

		chimeSdkIdentity_listAppInstanceAdminsCmd.Flags().String("app-instance-arn", "", "The ARN of the `AppInstance`.")
		chimeSdkIdentity_listAppInstanceAdminsCmd.Flags().String("max-results", "", "The maximum number of administrators that you want to return.")
		chimeSdkIdentity_listAppInstanceAdminsCmd.Flags().String("next-token", "", "The token returned from previous API requests until the number of administrators is reached.")
		chimeSdkIdentity_listAppInstanceAdminsCmd.MarkFlagRequired("app-instance-arn")
	})
	chimeSdkIdentityCmd.AddCommand(chimeSdkIdentity_listAppInstanceAdminsCmd)
}
