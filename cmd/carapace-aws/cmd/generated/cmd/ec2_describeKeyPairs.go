package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeKeyPairsCmd = &cobra.Command{
	Use:   "describe-key-pairs",
	Short: "Describes the specified key pairs or all of your key pairs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeKeyPairsCmd).Standalone()

	ec2_describeKeyPairsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeKeyPairsCmd.Flags().String("filters", "", "The filters.")
	ec2_describeKeyPairsCmd.Flags().Bool("include-public-key", false, "If `true`, the public key material is included in the response.")
	ec2_describeKeyPairsCmd.Flags().String("key-names", "", "The key pair names.")
	ec2_describeKeyPairsCmd.Flags().String("key-pair-ids", "", "The IDs of the key pairs.")
	ec2_describeKeyPairsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeKeyPairsCmd.Flags().Bool("no-include-public-key", false, "If `true`, the public key material is included in the response.")
	ec2_describeKeyPairsCmd.Flag("no-dry-run").Hidden = true
	ec2_describeKeyPairsCmd.Flag("no-include-public-key").Hidden = true
	ec2Cmd.AddCommand(ec2_describeKeyPairsCmd)
}
