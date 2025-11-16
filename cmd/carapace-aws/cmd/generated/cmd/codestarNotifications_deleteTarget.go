package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarNotifications_deleteTargetCmd = &cobra.Command{
	Use:   "delete-target",
	Short: "Deletes a specified target for notifications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarNotifications_deleteTargetCmd).Standalone()

	codestarNotifications_deleteTargetCmd.Flags().String("force-unsubscribe-all", "", "A Boolean value that can be used to delete all associations with this Amazon Q Developer in chat applications topic.")
	codestarNotifications_deleteTargetCmd.Flags().String("target-address", "", "The Amazon Resource Name (ARN) of the Amazon Q Developer in chat applications topic or Amazon Q Developer in chat applications client to delete.")
	codestarNotifications_deleteTargetCmd.MarkFlagRequired("target-address")
	codestarNotificationsCmd.AddCommand(codestarNotifications_deleteTargetCmd)
}
