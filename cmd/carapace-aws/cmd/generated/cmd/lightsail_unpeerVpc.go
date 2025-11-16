package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_unpeerVpcCmd = &cobra.Command{
	Use:   "unpeer-vpc",
	Short: "Unpeers the Lightsail VPC from the user's default VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_unpeerVpcCmd).Standalone()

	lightsailCmd.AddCommand(lightsail_unpeerVpcCmd)
}
