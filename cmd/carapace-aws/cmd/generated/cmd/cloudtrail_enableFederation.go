package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_enableFederationCmd = &cobra.Command{
	Use:   "enable-federation",
	Short: "Enables Lake query federation on the specified event data store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_enableFederationCmd).Standalone()

	cloudtrail_enableFederationCmd.Flags().String("event-data-store", "", "The ARN (or ID suffix of the ARN) of the event data store for which you want to enable Lake query federation.")
	cloudtrail_enableFederationCmd.Flags().String("federation-role-arn", "", "The ARN of the federation role to use for the event data store.")
	cloudtrail_enableFederationCmd.MarkFlagRequired("event-data-store")
	cloudtrail_enableFederationCmd.MarkFlagRequired("federation-role-arn")
	cloudtrailCmd.AddCommand(cloudtrail_enableFederationCmd)
}
