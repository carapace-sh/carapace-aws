package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyFpgaImageAttributeCmd = &cobra.Command{
	Use:   "modify-fpga-image-attribute",
	Short: "Modifies the specified attribute of the specified Amazon FPGA Image (AFI).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyFpgaImageAttributeCmd).Standalone()

	ec2_modifyFpgaImageAttributeCmd.Flags().String("attribute", "", "The name of the attribute.")
	ec2_modifyFpgaImageAttributeCmd.Flags().String("description", "", "A description for the AFI.")
	ec2_modifyFpgaImageAttributeCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyFpgaImageAttributeCmd.Flags().String("fpga-image-id", "", "The ID of the AFI.")
	ec2_modifyFpgaImageAttributeCmd.Flags().String("load-permission", "", "The load permission for the AFI.")
	ec2_modifyFpgaImageAttributeCmd.Flags().String("name", "", "A name for the AFI.")
	ec2_modifyFpgaImageAttributeCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyFpgaImageAttributeCmd.Flags().String("operation-type", "", "The operation type.")
	ec2_modifyFpgaImageAttributeCmd.Flags().String("product-codes", "", "The product codes.")
	ec2_modifyFpgaImageAttributeCmd.Flags().String("user-groups", "", "The user groups.")
	ec2_modifyFpgaImageAttributeCmd.Flags().String("user-ids", "", "The Amazon Web Services account IDs.")
	ec2_modifyFpgaImageAttributeCmd.MarkFlagRequired("fpga-image-id")
	ec2_modifyFpgaImageAttributeCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_modifyFpgaImageAttributeCmd)
}
