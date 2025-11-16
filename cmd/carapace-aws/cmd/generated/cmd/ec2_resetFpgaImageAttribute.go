package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_resetFpgaImageAttributeCmd = &cobra.Command{
	Use:   "reset-fpga-image-attribute",
	Short: "Resets the specified attribute of the specified Amazon FPGA Image (AFI) to its default value.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_resetFpgaImageAttributeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_resetFpgaImageAttributeCmd).Standalone()

		ec2_resetFpgaImageAttributeCmd.Flags().String("attribute", "", "The attribute.")
		ec2_resetFpgaImageAttributeCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_resetFpgaImageAttributeCmd.Flags().String("fpga-image-id", "", "The ID of the AFI.")
		ec2_resetFpgaImageAttributeCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_resetFpgaImageAttributeCmd.MarkFlagRequired("fpga-image-id")
		ec2_resetFpgaImageAttributeCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_resetFpgaImageAttributeCmd)
}
