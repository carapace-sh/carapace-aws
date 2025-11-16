package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_deleteStudioSessionMappingCmd = &cobra.Command{
	Use:   "delete-studio-session-mapping",
	Short: "Removes a user or group from an Amazon EMR Studio.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_deleteStudioSessionMappingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emr_deleteStudioSessionMappingCmd).Standalone()

		emr_deleteStudioSessionMappingCmd.Flags().String("identity-id", "", "The globally unique identifier (GUID) of the user or group to remove from the Amazon EMR Studio.")
		emr_deleteStudioSessionMappingCmd.Flags().String("identity-name", "", "The name of the user name or group to remove from the Amazon EMR Studio.")
		emr_deleteStudioSessionMappingCmd.Flags().String("identity-type", "", "Specifies whether the identity to delete from the Amazon EMR Studio is a user or a group.")
		emr_deleteStudioSessionMappingCmd.Flags().String("studio-id", "", "The ID of the Amazon EMR Studio.")
		emr_deleteStudioSessionMappingCmd.MarkFlagRequired("identity-type")
		emr_deleteStudioSessionMappingCmd.MarkFlagRequired("studio-id")
	})
	emrCmd.AddCommand(emr_deleteStudioSessionMappingCmd)
}
