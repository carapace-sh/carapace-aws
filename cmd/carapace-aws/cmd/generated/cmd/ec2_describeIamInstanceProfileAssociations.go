package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeIamInstanceProfileAssociationsCmd = &cobra.Command{
	Use:   "describe-iam-instance-profile-associations",
	Short: "Describes your IAM instance profile associations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeIamInstanceProfileAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeIamInstanceProfileAssociationsCmd).Standalone()

		ec2_describeIamInstanceProfileAssociationsCmd.Flags().String("association-ids", "", "The IAM instance profile associations.")
		ec2_describeIamInstanceProfileAssociationsCmd.Flags().String("filters", "", "The filters.")
		ec2_describeIamInstanceProfileAssociationsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_describeIamInstanceProfileAssociationsCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	})
	ec2Cmd.AddCommand(ec2_describeIamInstanceProfileAssociationsCmd)
}
