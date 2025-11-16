package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createFleetCmd = &cobra.Command{
	Use:   "create-fleet",
	Short: "Creates an EC2 Fleet that contains the configuration information for On-Demand Instances and Spot Instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createFleetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createFleetCmd).Standalone()

		ec2_createFleetCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ec2_createFleetCmd.Flags().String("context", "", "Reserved.")
		ec2_createFleetCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createFleetCmd.Flags().String("excess-capacity-termination-policy", "", "Indicates whether running instances should be terminated if the total target capacity of the EC2 Fleet is decreased below the current size of the EC2 Fleet.")
		ec2_createFleetCmd.Flags().String("launch-template-configs", "", "The configuration for the EC2 Fleet.")
		ec2_createFleetCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createFleetCmd.Flags().Bool("no-replace-unhealthy-instances", false, "Indicates whether EC2 Fleet should replace unhealthy Spot Instances.")
		ec2_createFleetCmd.Flags().Bool("no-terminate-instances-with-expiration", false, "Indicates whether running instances should be terminated when the EC2 Fleet expires.")
		ec2_createFleetCmd.Flags().String("on-demand-options", "", "Describes the configuration of On-Demand Instances in an EC2 Fleet.")
		ec2_createFleetCmd.Flags().Bool("replace-unhealthy-instances", false, "Indicates whether EC2 Fleet should replace unhealthy Spot Instances.")
		ec2_createFleetCmd.Flags().String("spot-options", "", "Describes the configuration of Spot Instances in an EC2 Fleet.")
		ec2_createFleetCmd.Flags().String("tag-specifications", "", "The key-value pair for tagging the EC2 Fleet request on creation.")
		ec2_createFleetCmd.Flags().String("target-capacity-specification", "", "The number of units to request.")
		ec2_createFleetCmd.Flags().Bool("terminate-instances-with-expiration", false, "Indicates whether running instances should be terminated when the EC2 Fleet expires.")
		ec2_createFleetCmd.Flags().String("type", "", "The fleet type.")
		ec2_createFleetCmd.Flags().String("valid-from", "", "The start date and time of the request, in UTC format (for example, *YYYY*-*MM*-*DD*T*HH*:*MM*:*SS*Z).")
		ec2_createFleetCmd.Flags().String("valid-until", "", "The end date and time of the request, in UTC format (for example, *YYYY*-*MM*-*DD*T*HH*:*MM*:*SS*Z).")
		ec2_createFleetCmd.MarkFlagRequired("launch-template-configs")
		ec2_createFleetCmd.Flag("no-dry-run").Hidden = true
		ec2_createFleetCmd.Flag("no-replace-unhealthy-instances").Hidden = true
		ec2_createFleetCmd.Flag("no-terminate-instances-with-expiration").Hidden = true
		ec2_createFleetCmd.MarkFlagRequired("target-capacity-specification")
	})
	ec2Cmd.AddCommand(ec2_createFleetCmd)
}
