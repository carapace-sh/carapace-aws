package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_peerVpcCmd = &cobra.Command{
	Use:   "peer-vpc",
	Short: "Peers the Lightsail VPC with the user's default VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_peerVpcCmd).Standalone()

	lightsailCmd.AddCommand(lightsail_peerVpcCmd)
}
