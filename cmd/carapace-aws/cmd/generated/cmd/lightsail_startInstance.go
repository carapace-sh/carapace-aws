package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_startInstanceCmd = &cobra.Command{
	Use:   "start-instance",
	Short: "Starts a specific Amazon Lightsail instance from a stopped state.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_startInstanceCmd).Standalone()

	lightsail_startInstanceCmd.Flags().String("instance-name", "", "The name of the instance (a virtual private server) to start.")
	lightsail_startInstanceCmd.MarkFlagRequired("instance-name")
	lightsailCmd.AddCommand(lightsail_startInstanceCmd)
}
