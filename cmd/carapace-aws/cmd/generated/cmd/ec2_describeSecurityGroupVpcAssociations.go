package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeSecurityGroupVpcAssociationsCmd = &cobra.Command{
	Use:   "describe-security-group-vpc-associations",
	Short: "Describes security group VPC associations made with [AssociateSecurityGroupVpc](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AssociateSecurityGroupVpc.html).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeSecurityGroupVpcAssociationsCmd).Standalone()

	ec2_describeSecurityGroupVpcAssociationsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeSecurityGroupVpcAssociationsCmd.Flags().String("filters", "", "Security group VPC association filters.")
	ec2_describeSecurityGroupVpcAssociationsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_describeSecurityGroupVpcAssociationsCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	ec2_describeSecurityGroupVpcAssociationsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeSecurityGroupVpcAssociationsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeSecurityGroupVpcAssociationsCmd)
}
