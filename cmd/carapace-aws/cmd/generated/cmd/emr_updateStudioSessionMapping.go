package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_updateStudioSessionMappingCmd = &cobra.Command{
	Use:   "update-studio-session-mapping",
	Short: "Updates the session policy attached to the user or group for the specified Amazon EMR Studio.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_updateStudioSessionMappingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emr_updateStudioSessionMappingCmd).Standalone()

		emr_updateStudioSessionMappingCmd.Flags().String("identity-id", "", "The globally unique identifier (GUID) of the user or group.")
		emr_updateStudioSessionMappingCmd.Flags().String("identity-name", "", "The name of the user or group to update.")
		emr_updateStudioSessionMappingCmd.Flags().String("identity-type", "", "Specifies whether the identity to update is a user or a group.")
		emr_updateStudioSessionMappingCmd.Flags().String("session-policy-arn", "", "The Amazon Resource Name (ARN) of the session policy to associate with the specified user or group.")
		emr_updateStudioSessionMappingCmd.Flags().String("studio-id", "", "The ID of the Amazon EMR Studio.")
		emr_updateStudioSessionMappingCmd.MarkFlagRequired("identity-type")
		emr_updateStudioSessionMappingCmd.MarkFlagRequired("session-policy-arn")
		emr_updateStudioSessionMappingCmd.MarkFlagRequired("studio-id")
	})
	emrCmd.AddCommand(emr_updateStudioSessionMappingCmd)
}
