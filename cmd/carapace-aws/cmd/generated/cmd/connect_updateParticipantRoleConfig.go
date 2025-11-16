package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateParticipantRoleConfigCmd = &cobra.Command{
	Use:   "update-participant-role-config",
	Short: "Updates timeouts for when human chat participants are to be considered idle, and when agents are automatically disconnected from a chat due to idleness.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateParticipantRoleConfigCmd).Standalone()

	connect_updateParticipantRoleConfigCmd.Flags().String("channel-configuration", "", "The Amazon Connect channel you want to configure.")
	connect_updateParticipantRoleConfigCmd.Flags().String("contact-id", "", "The identifier of the contact in this instance of Amazon Connect.")
	connect_updateParticipantRoleConfigCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_updateParticipantRoleConfigCmd.MarkFlagRequired("channel-configuration")
	connect_updateParticipantRoleConfigCmd.MarkFlagRequired("contact-id")
	connect_updateParticipantRoleConfigCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_updateParticipantRoleConfigCmd)
}
