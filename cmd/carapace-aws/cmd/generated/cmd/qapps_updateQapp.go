package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_updateQappCmd = &cobra.Command{
	Use:   "update-qapp",
	Short: "Updates an existing Amazon Q App, allowing modifications to its title, description, and definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_updateQappCmd).Standalone()

	qapps_updateQappCmd.Flags().String("app-definition", "", "The new definition specifying the cards and flow for the Q App.")
	qapps_updateQappCmd.Flags().String("app-id", "", "The unique identifier of the Q App to update.")
	qapps_updateQappCmd.Flags().String("description", "", "The new description for the Q App.")
	qapps_updateQappCmd.Flags().String("instance-id", "", "The unique identifier of the Amazon Q Business application environment instance.")
	qapps_updateQappCmd.Flags().String("title", "", "The new title for the Q App.")
	qapps_updateQappCmd.MarkFlagRequired("app-id")
	qapps_updateQappCmd.MarkFlagRequired("instance-id")
	qappsCmd.AddCommand(qapps_updateQappCmd)
}
