package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_pauseContactCmd = &cobra.Command{
	Use:   "pause-contact",
	Short: "Allows pausing an ongoing task contact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_pauseContactCmd).Standalone()

	connect_pauseContactCmd.Flags().String("contact-flow-id", "", "The identifier of the flow.")
	connect_pauseContactCmd.Flags().String("contact-id", "", "The identifier of the contact.")
	connect_pauseContactCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_pauseContactCmd.MarkFlagRequired("contact-id")
	connect_pauseContactCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_pauseContactCmd)
}
