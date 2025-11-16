package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_listStudioSessionMappingsCmd = &cobra.Command{
	Use:   "list-studio-session-mappings",
	Short: "Returns a list of all user or group session mappings for the Amazon EMR Studio specified by `StudioId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_listStudioSessionMappingsCmd).Standalone()

	emr_listStudioSessionMappingsCmd.Flags().String("identity-type", "", "Specifies whether to return session mappings for users or groups.")
	emr_listStudioSessionMappingsCmd.Flags().String("marker", "", "The pagination token that indicates the set of results to retrieve.")
	emr_listStudioSessionMappingsCmd.Flags().String("studio-id", "", "The ID of the Amazon EMR Studio.")
	emrCmd.AddCommand(emr_listStudioSessionMappingsCmd)
}
