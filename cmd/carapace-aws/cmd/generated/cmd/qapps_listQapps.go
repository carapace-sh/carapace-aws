package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qapps_listQappsCmd = &cobra.Command{
	Use:   "list-qapps",
	Short: "Lists the Amazon Q Apps owned by or associated with the user either because they created it or because they used it from the library in the past.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qapps_listQappsCmd).Standalone()

	qapps_listQappsCmd.Flags().String("instance-id", "", "The unique identifier of the Amazon Q Business application environment instance.")
	qapps_listQappsCmd.Flags().String("limit", "", "The maximum number of Q Apps to return in the response.")
	qapps_listQappsCmd.Flags().String("next-token", "", "The token to request the next page of results.")
	qapps_listQappsCmd.MarkFlagRequired("instance-id")
	qappsCmd.AddCommand(qapps_listQappsCmd)
}
