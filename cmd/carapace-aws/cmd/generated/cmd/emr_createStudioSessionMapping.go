package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_createStudioSessionMappingCmd = &cobra.Command{
	Use:   "create-studio-session-mapping",
	Short: "Maps a user or group to the Amazon EMR Studio specified by `StudioId`, and applies a session policy to refine Studio permissions for that user or group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_createStudioSessionMappingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emr_createStudioSessionMappingCmd).Standalone()

		emr_createStudioSessionMappingCmd.Flags().String("identity-id", "", "The globally unique identifier (GUID) of the user or group from the IAM Identity Center Identity Store.")
		emr_createStudioSessionMappingCmd.Flags().String("identity-name", "", "The name of the user or group.")
		emr_createStudioSessionMappingCmd.Flags().String("identity-type", "", "Specifies whether the identity to map to the Amazon EMR Studio is a user or a group.")
		emr_createStudioSessionMappingCmd.Flags().String("session-policy-arn", "", "The Amazon Resource Name (ARN) for the session policy that will be applied to the user or group.")
		emr_createStudioSessionMappingCmd.Flags().String("studio-id", "", "The ID of the Amazon EMR Studio to which the user or group will be mapped.")
		emr_createStudioSessionMappingCmd.MarkFlagRequired("identity-type")
		emr_createStudioSessionMappingCmd.MarkFlagRequired("session-policy-arn")
		emr_createStudioSessionMappingCmd.MarkFlagRequired("studio-id")
	})
	emrCmd.AddCommand(emr_createStudioSessionMappingCmd)
}
