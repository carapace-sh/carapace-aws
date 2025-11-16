package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_associateContactWithUserCmd = &cobra.Command{
	Use:   "associate-contact-with-user",
	Short: "Associates a queued contact with an agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_associateContactWithUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_associateContactWithUserCmd).Standalone()

		connect_associateContactWithUserCmd.Flags().String("contact-id", "", "The identifier of the contact in this instance of Amazon Connect.")
		connect_associateContactWithUserCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_associateContactWithUserCmd.Flags().String("user-id", "", "The identifier for the user.")
		connect_associateContactWithUserCmd.MarkFlagRequired("contact-id")
		connect_associateContactWithUserCmd.MarkFlagRequired("instance-id")
		connect_associateContactWithUserCmd.MarkFlagRequired("user-id")
	})
	connectCmd.AddCommand(connect_associateContactWithUserCmd)
}
