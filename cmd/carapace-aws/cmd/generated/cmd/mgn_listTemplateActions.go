package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_listTemplateActionsCmd = &cobra.Command{
	Use:   "list-template-actions",
	Short: "List template post migration custom actions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_listTemplateActionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgn_listTemplateActionsCmd).Standalone()

		mgn_listTemplateActionsCmd.Flags().String("filters", "", "Filters to apply when listing template post migration custom actions.")
		mgn_listTemplateActionsCmd.Flags().String("launch-configuration-template-id", "", "Launch configuration template ID.")
		mgn_listTemplateActionsCmd.Flags().String("max-results", "", "Maximum amount of items to return when listing template post migration custom actions.")
		mgn_listTemplateActionsCmd.Flags().String("next-token", "", "Next token to use when listing template post migration custom actions.")
		mgn_listTemplateActionsCmd.MarkFlagRequired("launch-configuration-template-id")
	})
	mgnCmd.AddCommand(mgn_listTemplateActionsCmd)
}
