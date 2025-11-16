package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_getTemplateSummaryCmd = &cobra.Command{
	Use:   "get-template-summary",
	Short: "Returns information about a new or existing template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_getTemplateSummaryCmd).Standalone()

	cloudformation_getTemplateSummaryCmd.Flags().String("call-as", "", "\\[Service-managed permissions] Specifies whether you are acting as an account administrator in the organization's management account or as a delegated administrator in a member account.")
	cloudformation_getTemplateSummaryCmd.Flags().String("stack-name", "", "The name or the stack ID that's associated with the stack, which aren't always interchangeable.")
	cloudformation_getTemplateSummaryCmd.Flags().String("stack-set-name", "", "The name or unique ID of the StackSet from which the stack was created.")
	cloudformation_getTemplateSummaryCmd.Flags().String("template-body", "", "Structure that contains the template body with a minimum length of 1 byte and a maximum length of 51,200 bytes.")
	cloudformation_getTemplateSummaryCmd.Flags().String("template-summary-config", "", "Specifies options for the `GetTemplateSummary` API action.")
	cloudformation_getTemplateSummaryCmd.Flags().String("template-url", "", "The URL of a file that contains the template body.")
	cloudformationCmd.AddCommand(cloudformation_getTemplateSummaryCmd)
}
