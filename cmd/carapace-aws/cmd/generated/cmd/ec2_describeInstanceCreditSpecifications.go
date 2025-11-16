package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeInstanceCreditSpecificationsCmd = &cobra.Command{
	Use:   "describe-instance-credit-specifications",
	Short: "Describes the credit option for CPU usage of the specified burstable performance instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeInstanceCreditSpecificationsCmd).Standalone()

	ec2_describeInstanceCreditSpecificationsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
	ec2_describeInstanceCreditSpecificationsCmd.Flags().String("filters", "", "The filters.")
	ec2_describeInstanceCreditSpecificationsCmd.Flags().String("instance-ids", "", "The instance IDs.")
	ec2_describeInstanceCreditSpecificationsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_describeInstanceCreditSpecificationsCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	ec2_describeInstanceCreditSpecificationsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
	ec2_describeInstanceCreditSpecificationsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeInstanceCreditSpecificationsCmd)
}
