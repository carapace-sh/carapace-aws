package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_isVpcPeeredCmd = &cobra.Command{
	Use:   "is-vpc-peered",
	Short: "Returns a Boolean value indicating whether your Lightsail VPC is peered.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_isVpcPeeredCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_isVpcPeeredCmd).Standalone()

	})
	lightsailCmd.AddCommand(lightsail_isVpcPeeredCmd)
}
