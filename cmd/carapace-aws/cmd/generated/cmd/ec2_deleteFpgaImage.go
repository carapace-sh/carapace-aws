package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteFpgaImageCmd = &cobra.Command{
	Use:   "delete-fpga-image",
	Short: "Deletes the specified Amazon FPGA Image (AFI).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteFpgaImageCmd).Standalone()

	ec2_deleteFpgaImageCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteFpgaImageCmd.Flags().String("fpga-image-id", "", "The ID of the AFI.")
	ec2_deleteFpgaImageCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteFpgaImageCmd.MarkFlagRequired("fpga-image-id")
	ec2_deleteFpgaImageCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_deleteFpgaImageCmd)
}
