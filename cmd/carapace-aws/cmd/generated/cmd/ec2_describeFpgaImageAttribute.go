package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeFpgaImageAttributeCmd = &cobra.Command{
	Use:   "describe-fpga-image-attribute",
	Short: "Describes the specified attribute of the specified Amazon FPGA Image (AFI).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeFpgaImageAttributeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeFpgaImageAttributeCmd).Standalone()

		ec2_describeFpgaImageAttributeCmd.Flags().String("attribute", "", "The AFI attribute.")
		ec2_describeFpgaImageAttributeCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeFpgaImageAttributeCmd.Flags().String("fpga-image-id", "", "The ID of the AFI.")
		ec2_describeFpgaImageAttributeCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeFpgaImageAttributeCmd.MarkFlagRequired("attribute")
		ec2_describeFpgaImageAttributeCmd.MarkFlagRequired("fpga-image-id")
		ec2_describeFpgaImageAttributeCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeFpgaImageAttributeCmd)
}
