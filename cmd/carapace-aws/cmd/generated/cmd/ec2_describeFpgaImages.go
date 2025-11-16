package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeFpgaImagesCmd = &cobra.Command{
	Use:   "describe-fpga-images",
	Short: "Describes the Amazon FPGA Images (AFIs) available to you.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeFpgaImagesCmd).Standalone()

	ec2_describeFpgaImagesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeFpgaImagesCmd.Flags().String("filters", "", "The filters.")
	ec2_describeFpgaImagesCmd.Flags().String("fpga-image-ids", "", "The AFI IDs.")
	ec2_describeFpgaImagesCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	ec2_describeFpgaImagesCmd.Flags().String("next-token", "", "The token to retrieve the next page of results.")
	ec2_describeFpgaImagesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeFpgaImagesCmd.Flags().String("owners", "", "Filters the AFI by owner.")
	ec2_describeFpgaImagesCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeFpgaImagesCmd)
}
