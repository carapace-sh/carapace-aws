package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_getStudioSessionMappingCmd = &cobra.Command{
	Use:   "get-studio-session-mapping",
	Short: "Fetches mapping details for the specified Amazon EMR Studio and identity (user or group).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_getStudioSessionMappingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emr_getStudioSessionMappingCmd).Standalone()

		emr_getStudioSessionMappingCmd.Flags().String("identity-id", "", "The globally unique identifier (GUID) of the user or group.")
		emr_getStudioSessionMappingCmd.Flags().String("identity-name", "", "The name of the user or group to fetch.")
		emr_getStudioSessionMappingCmd.Flags().String("identity-type", "", "Specifies whether the identity to fetch is a user or a group.")
		emr_getStudioSessionMappingCmd.Flags().String("studio-id", "", "The ID of the Amazon EMR Studio.")
		emr_getStudioSessionMappingCmd.MarkFlagRequired("identity-type")
		emr_getStudioSessionMappingCmd.MarkFlagRequired("studio-id")
	})
	emrCmd.AddCommand(emr_getStudioSessionMappingCmd)
}
