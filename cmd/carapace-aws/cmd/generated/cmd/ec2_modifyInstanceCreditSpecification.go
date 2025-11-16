package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyInstanceCreditSpecificationCmd = &cobra.Command{
	Use:   "modify-instance-credit-specification",
	Short: "Modifies the credit option for CPU usage on a running or stopped burstable performance instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyInstanceCreditSpecificationCmd).Standalone()

	ec2_modifyInstanceCreditSpecificationCmd.Flags().String("client-token", "", "A unique, case-sensitive token that you provide to ensure idempotency of your modification request.")
	ec2_modifyInstanceCreditSpecificationCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
	ec2_modifyInstanceCreditSpecificationCmd.Flags().String("instance-credit-specifications", "", "Information about the credit option for CPU usage.")
	ec2_modifyInstanceCreditSpecificationCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
	ec2_modifyInstanceCreditSpecificationCmd.MarkFlagRequired("instance-credit-specifications")
	ec2_modifyInstanceCreditSpecificationCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_modifyInstanceCreditSpecificationCmd)
}
