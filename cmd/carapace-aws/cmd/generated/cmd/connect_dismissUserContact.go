package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_dismissUserContactCmd = &cobra.Command{
	Use:   "dismiss-user-contact",
	Short: "Dismisses contacts from an agent’s CCP and returns the agent to an available state, which allows the agent to receive a new routed contact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_dismissUserContactCmd).Standalone()

	connect_dismissUserContactCmd.Flags().String("contact-id", "", "The identifier of the contact.")
	connect_dismissUserContactCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_dismissUserContactCmd.Flags().String("user-id", "", "The identifier of the user account.")
	connect_dismissUserContactCmd.MarkFlagRequired("contact-id")
	connect_dismissUserContactCmd.MarkFlagRequired("instance-id")
	connect_dismissUserContactCmd.MarkFlagRequired("user-id")
	connectCmd.AddCommand(connect_dismissUserContactCmd)
}
