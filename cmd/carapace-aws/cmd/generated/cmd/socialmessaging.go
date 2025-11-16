package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var socialmessagingCmd = &cobra.Command{
	Use:   "socialmessaging",
	Short: "*Amazon Web Services End User Messaging Social*, also referred to as Social messaging, is a messaging service that enables application developers to incorporate WhatsApp into their existing workflows.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(socialmessagingCmd).Standalone()

	rootCmd.AddCommand(socialmessagingCmd)
}
