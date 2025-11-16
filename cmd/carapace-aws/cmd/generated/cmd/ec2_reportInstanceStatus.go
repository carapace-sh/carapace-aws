package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_reportInstanceStatusCmd = &cobra.Command{
	Use:   "report-instance-status",
	Short: "Submits feedback about the status of an instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_reportInstanceStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_reportInstanceStatusCmd).Standalone()

		ec2_reportInstanceStatusCmd.Flags().String("description", "", "Descriptive text about the health state of your instance.")
		ec2_reportInstanceStatusCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_reportInstanceStatusCmd.Flags().String("end-time", "", "The time at which the reported instance health state ended.")
		ec2_reportInstanceStatusCmd.Flags().String("instances", "", "The instances.")
		ec2_reportInstanceStatusCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_reportInstanceStatusCmd.Flags().String("reason-codes", "", "The reason codes that describe the health state of your instance.")
		ec2_reportInstanceStatusCmd.Flags().String("start-time", "", "The time at which the reported instance health state began.")
		ec2_reportInstanceStatusCmd.Flags().String("status", "", "The status of all instances listed.")
		ec2_reportInstanceStatusCmd.MarkFlagRequired("instances")
		ec2_reportInstanceStatusCmd.Flag("no-dry-run").Hidden = true
		ec2_reportInstanceStatusCmd.MarkFlagRequired("reason-codes")
		ec2_reportInstanceStatusCmd.MarkFlagRequired("status")
	})
	ec2Cmd.AddCommand(ec2_reportInstanceStatusCmd)
}
