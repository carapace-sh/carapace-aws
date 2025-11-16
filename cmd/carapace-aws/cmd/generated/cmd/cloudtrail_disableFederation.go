package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_disableFederationCmd = &cobra.Command{
	Use:   "disable-federation",
	Short: "Disables Lake query federation on the specified event data store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_disableFederationCmd).Standalone()

	cloudtrail_disableFederationCmd.Flags().String("event-data-store", "", "The ARN (or ID suffix of the ARN) of the event data store for which you want to disable Lake query federation.")
	cloudtrail_disableFederationCmd.MarkFlagRequired("event-data-store")
	cloudtrailCmd.AddCommand(cloudtrail_disableFederationCmd)
}
