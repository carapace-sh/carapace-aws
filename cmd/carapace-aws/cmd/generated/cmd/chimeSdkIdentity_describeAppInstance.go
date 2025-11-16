package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkIdentity_describeAppInstanceCmd = &cobra.Command{
	Use:   "describe-app-instance",
	Short: "Returns the full details of an `AppInstance`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkIdentity_describeAppInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkIdentity_describeAppInstanceCmd).Standalone()

		chimeSdkIdentity_describeAppInstanceCmd.Flags().String("app-instance-arn", "", "The ARN of the `AppInstance`.")
		chimeSdkIdentity_describeAppInstanceCmd.MarkFlagRequired("app-instance-arn")
	})
	chimeSdkIdentityCmd.AddCommand(chimeSdkIdentity_describeAppInstanceCmd)
}
