package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_startResourceStateUpdateCmd = &cobra.Command{
	Use:   "start-resource-state-update",
	Short: "Begin asynchronous resource state update for lifecycle changes to the specified image resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_startResourceStateUpdateCmd).Standalone()

	imagebuilder_startResourceStateUpdateCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure idempotency of the request.")
	imagebuilder_startResourceStateUpdateCmd.Flags().String("exclusion-rules", "", "Skip action on the image resource and associated resources if specified exclusion rules are met.")
	imagebuilder_startResourceStateUpdateCmd.Flags().String("execution-role", "", "The name or Amazon Resource Name (ARN) of the IAM role that’s used to update image state.")
	imagebuilder_startResourceStateUpdateCmd.Flags().String("include-resources", "", "A list of image resources to update state for.")
	imagebuilder_startResourceStateUpdateCmd.Flags().String("resource-arn", "", "The ARN of the Image Builder resource that is updated.")
	imagebuilder_startResourceStateUpdateCmd.Flags().String("state", "", "Indicates the lifecycle action to take for this request.")
	imagebuilder_startResourceStateUpdateCmd.Flags().String("update-at", "", "The timestamp that indicates when resources are updated by a lifecycle action.")
	imagebuilder_startResourceStateUpdateCmd.MarkFlagRequired("client-token")
	imagebuilder_startResourceStateUpdateCmd.MarkFlagRequired("resource-arn")
	imagebuilder_startResourceStateUpdateCmd.MarkFlagRequired("state")
	imagebuilderCmd.AddCommand(imagebuilder_startResourceStateUpdateCmd)
}
