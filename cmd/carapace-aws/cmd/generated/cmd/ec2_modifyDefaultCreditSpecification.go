package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyDefaultCreditSpecificationCmd = &cobra.Command{
	Use:   "modify-default-credit-specification",
	Short: "Modifies the default credit option for CPU usage of burstable performance instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyDefaultCreditSpecificationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyDefaultCreditSpecificationCmd).Standalone()

		ec2_modifyDefaultCreditSpecificationCmd.Flags().String("cpu-credits", "", "The credit option for CPU usage of the instance family.")
		ec2_modifyDefaultCreditSpecificationCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_modifyDefaultCreditSpecificationCmd.Flags().String("instance-family", "", "The instance family.")
		ec2_modifyDefaultCreditSpecificationCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_modifyDefaultCreditSpecificationCmd.MarkFlagRequired("cpu-credits")
		ec2_modifyDefaultCreditSpecificationCmd.MarkFlagRequired("instance-family")
		ec2_modifyDefaultCreditSpecificationCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_modifyDefaultCreditSpecificationCmd)
}
