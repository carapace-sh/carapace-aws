package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createMacSystemIntegrityProtectionModificationTaskCmd = &cobra.Command{
	Use:   "create-mac-system-integrity-protection-modification-task",
	Short: "Creates a System Integrity Protection (SIP) modification task to configure the SIP settings for an x86 Mac instance or Apple silicon Mac instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createMacSystemIntegrityProtectionModificationTaskCmd).Standalone()

	ec2_createMacSystemIntegrityProtectionModificationTaskCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	ec2_createMacSystemIntegrityProtectionModificationTaskCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createMacSystemIntegrityProtectionModificationTaskCmd.Flags().String("instance-id", "", "The ID of the Amazon EC2 Mac instance.")
	ec2_createMacSystemIntegrityProtectionModificationTaskCmd.Flags().String("mac-credentials", "", "**\\[Apple silicon Mac instances only]** Specifies the following credentials:")
	ec2_createMacSystemIntegrityProtectionModificationTaskCmd.Flags().String("mac-system-integrity-protection-configuration", "", "Specifies the overrides to selectively enable or disable individual SIP settings.")
	ec2_createMacSystemIntegrityProtectionModificationTaskCmd.Flags().String("mac-system-integrity-protection-status", "", "Specifies the overall SIP status for the instance.")
	ec2_createMacSystemIntegrityProtectionModificationTaskCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createMacSystemIntegrityProtectionModificationTaskCmd.Flags().String("tag-specifications", "", "Specifies tags to apply to the SIP modification task.")
	ec2_createMacSystemIntegrityProtectionModificationTaskCmd.MarkFlagRequired("instance-id")
	ec2_createMacSystemIntegrityProtectionModificationTaskCmd.MarkFlagRequired("mac-system-integrity-protection-status")
	ec2_createMacSystemIntegrityProtectionModificationTaskCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_createMacSystemIntegrityProtectionModificationTaskCmd)
}
