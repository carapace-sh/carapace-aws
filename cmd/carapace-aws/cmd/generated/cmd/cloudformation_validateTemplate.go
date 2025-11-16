package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_validateTemplateCmd = &cobra.Command{
	Use:   "validate-template",
	Short: "Validates a specified template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_validateTemplateCmd).Standalone()

	cloudformation_validateTemplateCmd.Flags().String("template-body", "", "Structure that contains the template body with a minimum length of 1 byte and a maximum length of 51,200 bytes.")
	cloudformation_validateTemplateCmd.Flags().String("template-url", "", "The URL of a file that contains the template body.")
	cloudformationCmd.AddCommand(cloudformation_validateTemplateCmd)
}
