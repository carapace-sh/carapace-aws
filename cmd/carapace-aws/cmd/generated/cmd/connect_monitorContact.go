package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_monitorContactCmd = &cobra.Command{
	Use:   "monitor-contact",
	Short: "Initiates silent monitoring of a contact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_monitorContactCmd).Standalone()

	connect_monitorContactCmd.Flags().String("allowed-monitor-capabilities", "", "Specify which monitoring actions the user is allowed to take.")
	connect_monitorContactCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	connect_monitorContactCmd.Flags().String("contact-id", "", "The identifier of the contact.")
	connect_monitorContactCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_monitorContactCmd.Flags().String("user-id", "", "The identifier of the user account.")
	connect_monitorContactCmd.MarkFlagRequired("contact-id")
	connect_monitorContactCmd.MarkFlagRequired("instance-id")
	connect_monitorContactCmd.MarkFlagRequired("user-id")
	connectCmd.AddCommand(connect_monitorContactCmd)
}
