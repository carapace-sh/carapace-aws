package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_submitAttachmentStateChangesCmd = &cobra.Command{
	Use:   "submit-attachment-state-changes",
	Short: "This action is only used by the Amazon ECS agent, and it is not intended for use outside of the agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_submitAttachmentStateChangesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecs_submitAttachmentStateChangesCmd).Standalone()

		ecs_submitAttachmentStateChangesCmd.Flags().String("attachments", "", "Any attachments associated with the state change request.")
		ecs_submitAttachmentStateChangesCmd.Flags().String("cluster", "", "The short name or full ARN of the cluster that hosts the container instance the attachment belongs to.")
		ecs_submitAttachmentStateChangesCmd.MarkFlagRequired("attachments")
	})
	ecsCmd.AddCommand(ecs_submitAttachmentStateChangesCmd)
}
