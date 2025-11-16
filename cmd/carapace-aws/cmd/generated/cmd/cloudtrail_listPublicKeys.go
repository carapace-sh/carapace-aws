package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_listPublicKeysCmd = &cobra.Command{
	Use:   "list-public-keys",
	Short: "Returns all public keys whose private keys were used to sign the digest files within the specified time range.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_listPublicKeysCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudtrail_listPublicKeysCmd).Standalone()

		cloudtrail_listPublicKeysCmd.Flags().String("end-time", "", "Optionally specifies, in UTC, the end of the time range to look up public keys for CloudTrail digest files.")
		cloudtrail_listPublicKeysCmd.Flags().String("next-token", "", "Reserved for future use.")
		cloudtrail_listPublicKeysCmd.Flags().String("start-time", "", "Optionally specifies, in UTC, the start of the time range to look up public keys for CloudTrail digest files.")
	})
	cloudtrailCmd.AddCommand(cloudtrail_listPublicKeysCmd)
}
