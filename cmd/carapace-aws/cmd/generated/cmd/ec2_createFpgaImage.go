package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createFpgaImageCmd = &cobra.Command{
	Use:   "create-fpga-image",
	Short: "Creates an Amazon FPGA Image (AFI) from the specified design checkpoint (DCP).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createFpgaImageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createFpgaImageCmd).Standalone()

		ec2_createFpgaImageCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		ec2_createFpgaImageCmd.Flags().String("description", "", "A description for the AFI.")
		ec2_createFpgaImageCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createFpgaImageCmd.Flags().String("input-storage-location", "", "The location of the encrypted design checkpoint in Amazon S3.")
		ec2_createFpgaImageCmd.Flags().String("logs-storage-location", "", "The location in Amazon S3 for the output logs.")
		ec2_createFpgaImageCmd.Flags().String("name", "", "A name for the AFI.")
		ec2_createFpgaImageCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createFpgaImageCmd.Flags().String("tag-specifications", "", "The tags to apply to the FPGA image during creation.")
		ec2_createFpgaImageCmd.MarkFlagRequired("input-storage-location")
		ec2_createFpgaImageCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_createFpgaImageCmd)
}
