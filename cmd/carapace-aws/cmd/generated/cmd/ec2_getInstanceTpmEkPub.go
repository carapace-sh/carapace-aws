package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getInstanceTpmEkPubCmd = &cobra.Command{
	Use:   "get-instance-tpm-ek-pub",
	Short: "Gets the public endorsement key associated with the Nitro Trusted Platform Module (NitroTPM) for the specified instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getInstanceTpmEkPubCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getInstanceTpmEkPubCmd).Standalone()

		ec2_getInstanceTpmEkPubCmd.Flags().Bool("dry-run", false, "Specify this parameter to verify whether the request will succeed, without actually making the request.")
		ec2_getInstanceTpmEkPubCmd.Flags().String("instance-id", "", "The ID of the instance for which to get the public endorsement key.")
		ec2_getInstanceTpmEkPubCmd.Flags().String("key-format", "", "The required public endorsement key format.")
		ec2_getInstanceTpmEkPubCmd.Flags().String("key-type", "", "The required public endorsement key type.")
		ec2_getInstanceTpmEkPubCmd.Flags().Bool("no-dry-run", false, "Specify this parameter to verify whether the request will succeed, without actually making the request.")
		ec2_getInstanceTpmEkPubCmd.MarkFlagRequired("instance-id")
		ec2_getInstanceTpmEkPubCmd.MarkFlagRequired("key-format")
		ec2_getInstanceTpmEkPubCmd.MarkFlagRequired("key-type")
		ec2_getInstanceTpmEkPubCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_getInstanceTpmEkPubCmd)
}
