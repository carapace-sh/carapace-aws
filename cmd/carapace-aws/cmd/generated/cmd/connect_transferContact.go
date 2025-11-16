package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_transferContactCmd = &cobra.Command{
	Use:   "transfer-contact",
	Short: "Transfers `TASK` or `EMAIL` contacts from one agent or queue to another agent or queue at any point after a contact is created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_transferContactCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_transferContactCmd).Standalone()

		connect_transferContactCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		connect_transferContactCmd.Flags().String("contact-flow-id", "", "The identifier of the flow.")
		connect_transferContactCmd.Flags().String("contact-id", "", "The identifier of the contact in this instance of Amazon Connect.")
		connect_transferContactCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_transferContactCmd.Flags().String("queue-id", "", "The identifier for the queue.")
		connect_transferContactCmd.Flags().String("user-id", "", "The identifier for the user.")
		connect_transferContactCmd.MarkFlagRequired("contact-flow-id")
		connect_transferContactCmd.MarkFlagRequired("contact-id")
		connect_transferContactCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_transferContactCmd)
}
