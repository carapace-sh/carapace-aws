package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getInstanceMetadataDefaultsCmd = &cobra.Command{
	Use:   "get-instance-metadata-defaults",
	Short: "Gets the default instance metadata service (IMDS) settings that are set at the account level in the specified Amazon Web Services\u2028 Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getInstanceMetadataDefaultsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getInstanceMetadataDefaultsCmd).Standalone()

		ec2_getInstanceMetadataDefaultsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_getInstanceMetadataDefaultsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_getInstanceMetadataDefaultsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_getInstanceMetadataDefaultsCmd)
}
