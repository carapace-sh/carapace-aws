package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMessagingCmd = &cobra.Command{
	Use:   "chime-sdk-messaging",
	Short: "The Amazon Chime SDK messaging APIs in this section allow software developers to send and receive messages in custom messaging applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMessagingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMessagingCmd).Standalone()

	})
	rootCmd.AddCommand(chimeSdkMessagingCmd)
}
