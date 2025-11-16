package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_updateCollaborationCmd = &cobra.Command{
	Use:   "update-collaboration",
	Short: "Updates collaboration metadata and can only be called by the collaboration owner.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_updateCollaborationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_updateCollaborationCmd).Standalone()

		cleanrooms_updateCollaborationCmd.Flags().String("analytics-engine", "", "The analytics engine.")
		cleanrooms_updateCollaborationCmd.Flags().String("collaboration-identifier", "", "The identifier for the collaboration.")
		cleanrooms_updateCollaborationCmd.Flags().String("description", "", "A description of the collaboration.")
		cleanrooms_updateCollaborationCmd.Flags().String("name", "", "A human-readable identifier provided by the collaboration owner.")
		cleanrooms_updateCollaborationCmd.MarkFlagRequired("collaboration-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_updateCollaborationCmd)
}
