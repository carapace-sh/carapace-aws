package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_getQappSessionMetadataCmd = &cobra.Command{
	Use:   "get-qapp-session-metadata",
	Short: "Retrieves the current configuration of a Q App session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_getQappSessionMetadataCmd).Standalone()

	qapps_getQappSessionMetadataCmd.Flags().String("instance-id", "", "The unique identifier of the Amazon Q Business application environment instance.")
	qapps_getQappSessionMetadataCmd.Flags().String("session-id", "", "The unique identifier of the Q App session.")
	qapps_getQappSessionMetadataCmd.MarkFlagRequired("instance-id")
	qapps_getQappSessionMetadataCmd.MarkFlagRequired("session-id")
	qappsCmd.AddCommand(qapps_getQappSessionMetadataCmd)
}
