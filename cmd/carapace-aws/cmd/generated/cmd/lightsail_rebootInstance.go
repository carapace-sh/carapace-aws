package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_rebootInstanceCmd = &cobra.Command{
	Use:   "reboot-instance",
	Short: "Restarts a specific instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_rebootInstanceCmd).Standalone()

	lightsail_rebootInstanceCmd.Flags().String("instance-name", "", "The name of the instance to reboot.")
	lightsail_rebootInstanceCmd.MarkFlagRequired("instance-name")
	lightsailCmd.AddCommand(lightsail_rebootInstanceCmd)
}
