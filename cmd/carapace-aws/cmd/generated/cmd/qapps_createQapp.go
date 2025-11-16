package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_createQappCmd = &cobra.Command{
	Use:   "create-qapp",
	Short: "Creates a new Amazon Q App based on the provided definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_createQappCmd).Standalone()

	qapps_createQappCmd.Flags().String("app-definition", "", "The definition of the new Q App, specifying the cards and flow.")
	qapps_createQappCmd.Flags().String("description", "", "The description of the new Q App.")
	qapps_createQappCmd.Flags().String("instance-id", "", "The unique identifier of the Amazon Q Business application environment instance.")
	qapps_createQappCmd.Flags().String("tags", "", "Optional tags to associate with the new Q App.")
	qapps_createQappCmd.Flags().String("title", "", "The title of the new Q App.")
	qapps_createQappCmd.MarkFlagRequired("app-definition")
	qapps_createQappCmd.MarkFlagRequired("instance-id")
	qapps_createQappCmd.MarkFlagRequired("title")
	qappsCmd.AddCommand(qapps_createQappCmd)
}
