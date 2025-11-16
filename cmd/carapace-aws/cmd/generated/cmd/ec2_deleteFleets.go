package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteFleetsCmd = &cobra.Command{
	Use:   "delete-fleets",
	Short: "Deletes the specified EC2 Fleet request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteFleetsCmd).Standalone()

	ec2_deleteFleetsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteFleetsCmd.Flags().String("fleet-ids", "", "The IDs of the EC2 Fleets.")
	ec2_deleteFleetsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteFleetsCmd.Flags().Bool("no-terminate-instances", false, "Indicates whether to terminate the associated instances when the EC2 Fleet is deleted.")
	ec2_deleteFleetsCmd.Flags().Bool("terminate-instances", false, "Indicates whether to terminate the associated instances when the EC2 Fleet is deleted.")
	ec2_deleteFleetsCmd.MarkFlagRequired("fleet-ids")
	ec2_deleteFleetsCmd.Flag("no-dry-run").Hidden = true
	ec2_deleteFleetsCmd.Flag("no-terminate-instances").Hidden = true
	ec2_deleteFleetsCmd.MarkFlagRequired("no-terminate-instances")
	ec2_deleteFleetsCmd.MarkFlagRequired("terminate-instances")
	ec2Cmd.AddCommand(ec2_deleteFleetsCmd)
}
