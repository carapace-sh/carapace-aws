package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_detectStackResourceDriftCmd = &cobra.Command{
	Use:   "detect-stack-resource-drift",
	Short: "Returns information about whether a resource's actual configuration differs, or has *drifted*, from its expected configuration, as defined in the stack template and any values specified as template parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_detectStackResourceDriftCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_detectStackResourceDriftCmd).Standalone()

		cloudformation_detectStackResourceDriftCmd.Flags().String("logical-resource-id", "", "The logical name of the resource for which to return drift information.")
		cloudformation_detectStackResourceDriftCmd.Flags().String("stack-name", "", "The name of the stack to which the resource belongs.")
		cloudformation_detectStackResourceDriftCmd.MarkFlagRequired("logical-resource-id")
		cloudformation_detectStackResourceDriftCmd.MarkFlagRequired("stack-name")
	})
	cloudformationCmd.AddCommand(cloudformation_detectStackResourceDriftCmd)
}
