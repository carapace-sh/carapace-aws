package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkIdentity_describeAppInstanceUserCmd = &cobra.Command{
	Use:   "describe-app-instance-user",
	Short: "Returns the full details of an `AppInstanceUser`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkIdentity_describeAppInstanceUserCmd).Standalone()

	chimeSdkIdentity_describeAppInstanceUserCmd.Flags().String("app-instance-user-arn", "", "The ARN of the `AppInstanceUser`.")
	chimeSdkIdentity_describeAppInstanceUserCmd.MarkFlagRequired("app-instance-user-arn")
	chimeSdkIdentityCmd.AddCommand(chimeSdkIdentity_describeAppInstanceUserCmd)
}
