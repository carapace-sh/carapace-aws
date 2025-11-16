package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyInstanceMetadataDefaultsCmd = &cobra.Command{
	Use:   "modify-instance-metadata-defaults",
	Short: "Modifies the default instance metadata service (IMDS) settings at the account level in the specified Amazon Web Services\u2028 Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyInstanceMetadataDefaultsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyInstanceMetadataDefaultsCmd).Standalone()

		ec2_modifyInstanceMetadataDefaultsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_modifyInstanceMetadataDefaultsCmd.Flags().String("http-endpoint", "", "Enables or disables the IMDS endpoint on an instance.")
		ec2_modifyInstanceMetadataDefaultsCmd.Flags().String("http-put-response-hop-limit", "", "The maximum number of hops that the metadata token can travel.")
		ec2_modifyInstanceMetadataDefaultsCmd.Flags().String("http-tokens", "", "Indicates whether IMDSv2 is required.")
		ec2_modifyInstanceMetadataDefaultsCmd.Flags().String("instance-metadata-tags", "", "Enables or disables access to an instance's tags from the instance metadata.")
		ec2_modifyInstanceMetadataDefaultsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
		ec2_modifyInstanceMetadataDefaultsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_modifyInstanceMetadataDefaultsCmd)
}
