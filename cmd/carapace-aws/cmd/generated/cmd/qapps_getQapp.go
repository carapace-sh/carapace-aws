package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_getQappCmd = &cobra.Command{
	Use:   "get-qapp",
	Short: "Retrieves the full details of an Q App, including its definition specifying the cards and flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_getQappCmd).Standalone()

	qapps_getQappCmd.Flags().String("app-id", "", "The unique identifier of the Q App to retrieve.")
	qapps_getQappCmd.Flags().String("app-version", "", "The version of the Q App.")
	qapps_getQappCmd.Flags().String("instance-id", "", "The unique identifier of the Amazon Q Business application environment instance.")
	qapps_getQappCmd.MarkFlagRequired("app-id")
	qapps_getQappCmd.MarkFlagRequired("instance-id")
	qappsCmd.AddCommand(qapps_getQappCmd)
}
