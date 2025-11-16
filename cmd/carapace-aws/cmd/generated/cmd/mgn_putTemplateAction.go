package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_putTemplateActionCmd = &cobra.Command{
	Use:   "put-template-action",
	Short: "Put template post migration custom action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_putTemplateActionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgn_putTemplateActionCmd).Standalone()

		mgn_putTemplateActionCmd.Flags().String("action-id", "", "Template post migration custom action ID.")
		mgn_putTemplateActionCmd.Flags().String("action-name", "", "Template post migration custom action name.")
		mgn_putTemplateActionCmd.Flags().Bool("active", false, "Template post migration custom action active status.")
		mgn_putTemplateActionCmd.Flags().String("category", "", "Template post migration custom action category.")
		mgn_putTemplateActionCmd.Flags().String("description", "", "Template post migration custom action description.")
		mgn_putTemplateActionCmd.Flags().String("document-identifier", "", "Template post migration custom action document identifier.")
		mgn_putTemplateActionCmd.Flags().String("document-version", "", "Template post migration custom action document version.")
		mgn_putTemplateActionCmd.Flags().String("external-parameters", "", "Template post migration custom action external parameters.")
		mgn_putTemplateActionCmd.Flags().String("launch-configuration-template-id", "", "Launch configuration template ID.")
		mgn_putTemplateActionCmd.Flags().Bool("must-succeed-for-cutover", false, "Template post migration custom action must succeed for cutover.")
		mgn_putTemplateActionCmd.Flags().Bool("no-active", false, "Template post migration custom action active status.")
		mgn_putTemplateActionCmd.Flags().Bool("no-must-succeed-for-cutover", false, "Template post migration custom action must succeed for cutover.")
		mgn_putTemplateActionCmd.Flags().String("operating-system", "", "Operating system eligible for this template post migration custom action.")
		mgn_putTemplateActionCmd.Flags().String("order", "", "Template post migration custom action order.")
		mgn_putTemplateActionCmd.Flags().String("parameters", "", "Template post migration custom action parameters.")
		mgn_putTemplateActionCmd.Flags().String("timeout-seconds", "", "Template post migration custom action timeout in seconds.")
		mgn_putTemplateActionCmd.MarkFlagRequired("action-id")
		mgn_putTemplateActionCmd.MarkFlagRequired("action-name")
		mgn_putTemplateActionCmd.MarkFlagRequired("document-identifier")
		mgn_putTemplateActionCmd.MarkFlagRequired("launch-configuration-template-id")
		mgn_putTemplateActionCmd.Flag("no-active").Hidden = true
		mgn_putTemplateActionCmd.Flag("no-must-succeed-for-cutover").Hidden = true
		mgn_putTemplateActionCmd.MarkFlagRequired("order")
	})
	mgnCmd.AddCommand(mgn_putTemplateActionCmd)
}
