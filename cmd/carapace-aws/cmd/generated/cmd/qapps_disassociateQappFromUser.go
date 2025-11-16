package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_disassociateQappFromUserCmd = &cobra.Command{
	Use:   "disassociate-qapp-from-user",
	Short: "Disassociates a Q App from a user removing the user's access to run the Q App.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_disassociateQappFromUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qapps_disassociateQappFromUserCmd).Standalone()

		qapps_disassociateQappFromUserCmd.Flags().String("app-id", "", "The unique identifier of the Q App to disassociate from the user.")
		qapps_disassociateQappFromUserCmd.Flags().String("instance-id", "", "The unique identifier of the Amazon Q Business application environment instance.")
		qapps_disassociateQappFromUserCmd.MarkFlagRequired("app-id")
		qapps_disassociateQappFromUserCmd.MarkFlagRequired("instance-id")
	})
	qappsCmd.AddCommand(qapps_disassociateQappFromUserCmd)
}
