package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_copyFpgaImageCmd = &cobra.Command{
	Use:   "copy-fpga-image",
	Short: "Copies the specified Amazon FPGA Image (AFI) to the current Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_copyFpgaImageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_copyFpgaImageCmd).Standalone()

		ec2_copyFpgaImageCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ec2_copyFpgaImageCmd.Flags().String("description", "", "The description for the new AFI.")
		ec2_copyFpgaImageCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_copyFpgaImageCmd.Flags().String("name", "", "The name for the new AFI.")
		ec2_copyFpgaImageCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_copyFpgaImageCmd.Flags().String("source-fpga-image-id", "", "The ID of the source AFI.")
		ec2_copyFpgaImageCmd.Flags().String("source-region", "", "The Region that contains the source AFI.")
		ec2_copyFpgaImageCmd.Flag("no-dry-run").Hidden = true
		ec2_copyFpgaImageCmd.MarkFlagRequired("source-fpga-image-id")
		ec2_copyFpgaImageCmd.MarkFlagRequired("source-region")
	})
	ec2Cmd.AddCommand(ec2_copyFpgaImageCmd)
}
