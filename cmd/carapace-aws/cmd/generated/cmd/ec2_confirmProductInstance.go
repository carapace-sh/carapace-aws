package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_confirmProductInstanceCmd = &cobra.Command{
	Use:   "confirm-product-instance",
	Short: "Determines whether a product code is associated with an instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_confirmProductInstanceCmd).Standalone()

	ec2_confirmProductInstanceCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
	ec2_confirmProductInstanceCmd.Flags().String("instance-id", "", "The ID of the instance.")
	ec2_confirmProductInstanceCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
	ec2_confirmProductInstanceCmd.Flags().String("product-code", "", "The product code.")
	ec2_confirmProductInstanceCmd.MarkFlagRequired("instance-id")
	ec2_confirmProductInstanceCmd.Flag("no-dry-run").Hidden = true
	ec2_confirmProductInstanceCmd.MarkFlagRequired("product-code")
	ec2Cmd.AddCommand(ec2_confirmProductInstanceCmd)
}
