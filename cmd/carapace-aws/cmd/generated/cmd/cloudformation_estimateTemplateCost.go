package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_estimateTemplateCostCmd = &cobra.Command{
	Use:   "estimate-template-cost",
	Short: "Returns the estimated monthly cost of a template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_estimateTemplateCostCmd).Standalone()

	cloudformation_estimateTemplateCostCmd.Flags().String("parameters", "", "A list of `Parameter` structures that specify input parameters.")
	cloudformation_estimateTemplateCostCmd.Flags().String("template-body", "", "Structure that contains the template body with a minimum length of 1 byte and a maximum length of 51,200 bytes.")
	cloudformation_estimateTemplateCostCmd.Flags().String("template-url", "", "The URL of a file that contains the template body.")
	cloudformationCmd.AddCommand(cloudformation_estimateTemplateCostCmd)
}
