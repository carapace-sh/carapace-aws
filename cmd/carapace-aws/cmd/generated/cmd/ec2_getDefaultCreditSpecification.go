package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getDefaultCreditSpecificationCmd = &cobra.Command{
	Use:   "get-default-credit-specification",
	Short: "Describes the default credit option for CPU usage of a burstable performance instance family.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getDefaultCreditSpecificationCmd).Standalone()

	ec2_getDefaultCreditSpecificationCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
	ec2_getDefaultCreditSpecificationCmd.Flags().String("instance-family", "", "The instance family.")
	ec2_getDefaultCreditSpecificationCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
	ec2_getDefaultCreditSpecificationCmd.MarkFlagRequired("instance-family")
	ec2_getDefaultCreditSpecificationCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_getDefaultCreditSpecificationCmd)
}
