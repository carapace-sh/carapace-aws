package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_detectStackDriftCmd = &cobra.Command{
	Use:   "detect-stack-drift",
	Short: "Detects whether a stack's actual configuration differs, or has *drifted*, from its expected configuration, as defined in the stack template and any values specified as template parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_detectStackDriftCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_detectStackDriftCmd).Standalone()

		cloudformation_detectStackDriftCmd.Flags().String("logical-resource-ids", "", "The logical names of any resources you want to use as filters.")
		cloudformation_detectStackDriftCmd.Flags().String("stack-name", "", "The name of the stack for which you want to detect drift.")
		cloudformation_detectStackDriftCmd.MarkFlagRequired("stack-name")
	})
	cloudformationCmd.AddCommand(cloudformation_detectStackDriftCmd)
}
