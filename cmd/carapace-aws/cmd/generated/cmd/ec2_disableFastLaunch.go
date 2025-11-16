package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disableFastLaunchCmd = &cobra.Command{
	Use:   "disable-fast-launch",
	Short: "Discontinue Windows fast launch for a Windows AMI, and clean up existing pre-provisioned snapshots.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disableFastLaunchCmd).Standalone()

	ec2_disableFastLaunchCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_disableFastLaunchCmd.Flags().Bool("force", false, "Forces the image settings to turn off Windows fast launch for your Windows AMI.")
	ec2_disableFastLaunchCmd.Flags().String("image-id", "", "Specify the ID of the image for which to disable Windows fast launch.")
	ec2_disableFastLaunchCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_disableFastLaunchCmd.Flags().Bool("no-force", false, "Forces the image settings to turn off Windows fast launch for your Windows AMI.")
	ec2_disableFastLaunchCmd.MarkFlagRequired("image-id")
	ec2_disableFastLaunchCmd.Flag("no-dry-run").Hidden = true
	ec2_disableFastLaunchCmd.Flag("no-force").Hidden = true
	ec2Cmd.AddCommand(ec2_disableFastLaunchCmd)
}
