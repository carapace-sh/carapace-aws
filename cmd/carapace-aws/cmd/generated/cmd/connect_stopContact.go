package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_stopContactCmd = &cobra.Command{
	Use:   "stop-contact",
	Short: "Ends the specified contact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_stopContactCmd).Standalone()

	connect_stopContactCmd.Flags().String("contact-id", "", "The ID of the contact.")
	connect_stopContactCmd.Flags().String("disconnect-reason", "", "The reason a contact can be disconnected.")
	connect_stopContactCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_stopContactCmd.MarkFlagRequired("contact-id")
	connect_stopContactCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_stopContactCmd)
}
