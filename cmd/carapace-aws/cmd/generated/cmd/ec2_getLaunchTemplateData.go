package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getLaunchTemplateDataCmd = &cobra.Command{
	Use:   "get-launch-template-data",
	Short: "Retrieves the configuration data of the specified instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getLaunchTemplateDataCmd).Standalone()

	ec2_getLaunchTemplateDataCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_getLaunchTemplateDataCmd.Flags().String("instance-id", "", "The ID of the instance.")
	ec2_getLaunchTemplateDataCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_getLaunchTemplateDataCmd.MarkFlagRequired("instance-id")
	ec2_getLaunchTemplateDataCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_getLaunchTemplateDataCmd)
}
