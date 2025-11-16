package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_updateStandardsControlCmd = &cobra.Command{
	Use:   "update-standards-control",
	Short: "Used to control whether an individual security standard control is enabled or disabled.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_updateStandardsControlCmd).Standalone()

	securityhub_updateStandardsControlCmd.Flags().String("control-status", "", "The updated status of the security standard control.")
	securityhub_updateStandardsControlCmd.Flags().String("disabled-reason", "", "A description of the reason why you are disabling a security standard control.")
	securityhub_updateStandardsControlCmd.Flags().String("standards-control-arn", "", "The ARN of the security standard control to enable or disable.")
	securityhub_updateStandardsControlCmd.MarkFlagRequired("standards-control-arn")
	securityhubCmd.AddCommand(securityhub_updateStandardsControlCmd)
}
