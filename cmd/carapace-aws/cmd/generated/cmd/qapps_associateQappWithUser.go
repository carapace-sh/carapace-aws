package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_associateQappWithUserCmd = &cobra.Command{
	Use:   "associate-qapp-with-user",
	Short: "This operation creates a link between the user's identity calling the operation and a specific Q App.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_associateQappWithUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qapps_associateQappWithUserCmd).Standalone()

		qapps_associateQappWithUserCmd.Flags().String("app-id", "", "The ID of the Amazon Q App to associate with the user.")
		qapps_associateQappWithUserCmd.Flags().String("instance-id", "", "The unique identifier of the Amazon Q Business application environment instance.")
		qapps_associateQappWithUserCmd.MarkFlagRequired("app-id")
		qapps_associateQappWithUserCmd.MarkFlagRequired("instance-id")
	})
	qappsCmd.AddCommand(qapps_associateQappWithUserCmd)
}
