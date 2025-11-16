package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_deleteQappCmd = &cobra.Command{
	Use:   "delete-qapp",
	Short: "Deletes an Amazon Q App owned by the user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_deleteQappCmd).Standalone()

	qapps_deleteQappCmd.Flags().String("app-id", "", "The unique identifier of the Q App to delete.")
	qapps_deleteQappCmd.Flags().String("instance-id", "", "The unique identifier of the Amazon Q Business application environment instance.")
	qapps_deleteQappCmd.MarkFlagRequired("app-id")
	qapps_deleteQappCmd.MarkFlagRequired("instance-id")
	qappsCmd.AddCommand(qapps_deleteQappCmd)
}
