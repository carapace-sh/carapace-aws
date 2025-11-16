package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_createEnvironmentTemplateCmd = &cobra.Command{
	Use:   "create-environment-template",
	Short: "Create an environment template for Proton.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_createEnvironmentTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_createEnvironmentTemplateCmd).Standalone()

		proton_createEnvironmentTemplateCmd.Flags().String("description", "", "A description of the environment template.")
		proton_createEnvironmentTemplateCmd.Flags().String("display-name", "", "The environment template name as displayed in the developer interface.")
		proton_createEnvironmentTemplateCmd.Flags().String("encryption-key", "", "A customer provided encryption key that Proton uses to encrypt data.")
		proton_createEnvironmentTemplateCmd.Flags().String("name", "", "The name of the environment template.")
		proton_createEnvironmentTemplateCmd.Flags().String("provisioning", "", "When included, indicates that the environment template is for customer provisioned and managed infrastructure.")
		proton_createEnvironmentTemplateCmd.Flags().String("tags", "", "An optional list of metadata items that you can associate with the Proton environment template.")
		proton_createEnvironmentTemplateCmd.MarkFlagRequired("name")
	})
	protonCmd.AddCommand(proton_createEnvironmentTemplateCmd)
}
