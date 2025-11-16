package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_deleteDistributionCmd = &cobra.Command{
	Use:   "delete-distribution",
	Short: "Deletes your Amazon Lightsail content delivery network (CDN) distribution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_deleteDistributionCmd).Standalone()

	lightsail_deleteDistributionCmd.Flags().String("distribution-name", "", "The name of the distribution to delete.")
	lightsailCmd.AddCommand(lightsail_deleteDistributionCmd)
}
