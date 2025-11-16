package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getRegionsCmd = &cobra.Command{
	Use:   "get-regions",
	Short: "Returns a list of all valid regions for Amazon Lightsail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getRegionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_getRegionsCmd).Standalone()

		lightsail_getRegionsCmd.Flags().String("include-availability-zones", "", "A Boolean value indicating whether to also include Availability Zones in your get regions request.")
		lightsail_getRegionsCmd.Flags().String("include-relational-database-availability-zones", "", "A Boolean value indicating whether to also include Availability Zones for databases in your get regions request.")
	})
	lightsailCmd.AddCommand(lightsail_getRegionsCmd)
}
