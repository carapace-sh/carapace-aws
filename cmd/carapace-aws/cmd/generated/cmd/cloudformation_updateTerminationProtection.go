package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_updateTerminationProtectionCmd = &cobra.Command{
	Use:   "update-termination-protection",
	Short: "Updates termination protection for the specified stack.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_updateTerminationProtectionCmd).Standalone()

	cloudformation_updateTerminationProtectionCmd.Flags().String("enable-termination-protection", "", "Whether to enable termination protection on the specified stack.")
	cloudformation_updateTerminationProtectionCmd.Flags().String("stack-name", "", "The name or unique ID of the stack for which you want to set termination protection.")
	cloudformation_updateTerminationProtectionCmd.MarkFlagRequired("enable-termination-protection")
	cloudformation_updateTerminationProtectionCmd.MarkFlagRequired("stack-name")
	cloudformationCmd.AddCommand(cloudformation_updateTerminationProtectionCmd)
}
