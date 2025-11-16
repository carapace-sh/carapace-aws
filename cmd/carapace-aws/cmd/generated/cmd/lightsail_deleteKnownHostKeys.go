package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_deleteKnownHostKeysCmd = &cobra.Command{
	Use:   "delete-known-host-keys",
	Short: "Deletes the known host key or certificate used by the Amazon Lightsail browser-based SSH or RDP clients to authenticate an instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_deleteKnownHostKeysCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_deleteKnownHostKeysCmd).Standalone()

		lightsail_deleteKnownHostKeysCmd.Flags().String("instance-name", "", "The name of the instance for which you want to reset the host key or certificate.")
		lightsail_deleteKnownHostKeysCmd.MarkFlagRequired("instance-name")
	})
	lightsailCmd.AddCommand(lightsail_deleteKnownHostKeysCmd)
}
