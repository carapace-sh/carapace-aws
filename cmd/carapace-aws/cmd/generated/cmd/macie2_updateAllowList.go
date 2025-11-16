package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_updateAllowListCmd = &cobra.Command{
	Use:   "update-allow-list",
	Short: "Updates the settings for an allow list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_updateAllowListCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_updateAllowListCmd).Standalone()

		macie2_updateAllowListCmd.Flags().String("criteria", "", "The criteria that specify the text or text pattern to ignore.")
		macie2_updateAllowListCmd.Flags().String("description", "", "A custom description of the allow list.")
		macie2_updateAllowListCmd.Flags().String("id", "", "The unique identifier for the Amazon Macie resource that the request applies to.")
		macie2_updateAllowListCmd.Flags().String("name", "", "A custom name for the allow list.")
		macie2_updateAllowListCmd.MarkFlagRequired("criteria")
		macie2_updateAllowListCmd.MarkFlagRequired("id")
		macie2_updateAllowListCmd.MarkFlagRequired("name")
	})
	macie2Cmd.AddCommand(macie2_updateAllowListCmd)
}
