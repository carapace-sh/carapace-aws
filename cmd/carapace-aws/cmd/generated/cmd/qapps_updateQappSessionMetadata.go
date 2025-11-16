package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_updateQappSessionMetadataCmd = &cobra.Command{
	Use:   "update-qapp-session-metadata",
	Short: "Updates the configuration metadata of a session for a given Q App `sessionId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_updateQappSessionMetadataCmd).Standalone()

	qapps_updateQappSessionMetadataCmd.Flags().String("instance-id", "", "The unique identifier of the Amazon Q Business application environment instance.")
	qapps_updateQappSessionMetadataCmd.Flags().String("session-id", "", "The unique identifier of the Q App session to update configuration for.")
	qapps_updateQappSessionMetadataCmd.Flags().String("session-name", "", "The new name for the Q App session.")
	qapps_updateQappSessionMetadataCmd.Flags().String("sharing-configuration", "", "The new sharing configuration for the Q App data collection session.")
	qapps_updateQappSessionMetadataCmd.MarkFlagRequired("instance-id")
	qapps_updateQappSessionMetadataCmd.MarkFlagRequired("session-id")
	qapps_updateQappSessionMetadataCmd.MarkFlagRequired("sharing-configuration")
	qappsCmd.AddCommand(qapps_updateQappSessionMetadataCmd)
}
