package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_describeStackResourceCmd = &cobra.Command{
	Use:   "describe-stack-resource",
	Short: "Returns a description of the specified resource in the specified stack.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_describeStackResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_describeStackResourceCmd).Standalone()

		cloudformation_describeStackResourceCmd.Flags().String("logical-resource-id", "", "The logical name of the resource as specified in the template.")
		cloudformation_describeStackResourceCmd.Flags().String("stack-name", "", "The name or the unique stack ID that's associated with the stack, which aren't always interchangeable:")
		cloudformation_describeStackResourceCmd.MarkFlagRequired("logical-resource-id")
		cloudformation_describeStackResourceCmd.MarkFlagRequired("stack-name")
	})
	cloudformationCmd.AddCommand(cloudformation_describeStackResourceCmd)
}
