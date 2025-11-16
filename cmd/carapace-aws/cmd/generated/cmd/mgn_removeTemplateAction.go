package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_removeTemplateActionCmd = &cobra.Command{
	Use:   "remove-template-action",
	Short: "Remove template post migration custom action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_removeTemplateActionCmd).Standalone()

	mgn_removeTemplateActionCmd.Flags().String("action-id", "", "Template post migration custom action ID to remove.")
	mgn_removeTemplateActionCmd.Flags().String("launch-configuration-template-id", "", "Launch configuration template ID of the post migration custom action to remove.")
	mgn_removeTemplateActionCmd.MarkFlagRequired("action-id")
	mgn_removeTemplateActionCmd.MarkFlagRequired("launch-configuration-template-id")
	mgnCmd.AddCommand(mgn_removeTemplateActionCmd)
}
