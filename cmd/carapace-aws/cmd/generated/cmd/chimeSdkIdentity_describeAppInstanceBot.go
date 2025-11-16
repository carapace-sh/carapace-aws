package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkIdentity_describeAppInstanceBotCmd = &cobra.Command{
	Use:   "describe-app-instance-bot",
	Short: "The `AppInstanceBot's` information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkIdentity_describeAppInstanceBotCmd).Standalone()

	chimeSdkIdentity_describeAppInstanceBotCmd.Flags().String("app-instance-bot-arn", "", "The ARN of the `AppInstanceBot`.")
	chimeSdkIdentity_describeAppInstanceBotCmd.MarkFlagRequired("app-instance-bot-arn")
	chimeSdkIdentityCmd.AddCommand(chimeSdkIdentity_describeAppInstanceBotCmd)
}
