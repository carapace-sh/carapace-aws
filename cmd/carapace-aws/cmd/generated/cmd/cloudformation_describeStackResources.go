package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_describeStackResourcesCmd = &cobra.Command{
	Use:   "describe-stack-resources",
	Short: "Returns Amazon Web Services resource descriptions for running and deleted stacks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_describeStackResourcesCmd).Standalone()

	cloudformation_describeStackResourcesCmd.Flags().String("logical-resource-id", "", "The logical name of the resource as specified in the template.")
	cloudformation_describeStackResourcesCmd.Flags().String("physical-resource-id", "", "The name or unique identifier that corresponds to a physical instance ID of a resource supported by CloudFormation.")
	cloudformation_describeStackResourcesCmd.Flags().String("stack-name", "", "The name or the unique stack ID that is associated with the stack, which aren't always interchangeable:")
	cloudformationCmd.AddCommand(cloudformation_describeStackResourcesCmd)
}
