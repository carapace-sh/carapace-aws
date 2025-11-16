package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_listScenesCmd = &cobra.Command{
	Use:   "list-scenes",
	Short: "Lists all scenes in a workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_listScenesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iottwinmaker_listScenesCmd).Standalone()

		iottwinmaker_listScenesCmd.Flags().String("max-results", "", "Specifies the maximum number of results to display.")
		iottwinmaker_listScenesCmd.Flags().String("next-token", "", "The string that specifies the next page of results.")
		iottwinmaker_listScenesCmd.Flags().String("workspace-id", "", "The ID of the workspace that contains the scenes.")
		iottwinmaker_listScenesCmd.MarkFlagRequired("workspace-id")
	})
	iottwinmakerCmd.AddCommand(iottwinmaker_listScenesCmd)
}
