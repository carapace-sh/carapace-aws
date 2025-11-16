package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wellarchitected_updateShareInvitationCmd = &cobra.Command{
	Use:   "update-share-invitation",
	Short: "Update a workload or custom lens share invitation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wellarchitected_updateShareInvitationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wellarchitected_updateShareInvitationCmd).Standalone()

		wellarchitected_updateShareInvitationCmd.Flags().String("share-invitation-action", "", "")
		wellarchitected_updateShareInvitationCmd.Flags().String("share-invitation-id", "", "The ID assigned to the share invitation.")
		wellarchitected_updateShareInvitationCmd.MarkFlagRequired("share-invitation-action")
		wellarchitected_updateShareInvitationCmd.MarkFlagRequired("share-invitation-id")
	})
	wellarchitectedCmd.AddCommand(wellarchitected_updateShareInvitationCmd)
}
