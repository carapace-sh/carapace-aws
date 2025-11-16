package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_describeRoutingProfileCmd = &cobra.Command{
	Use:   "describe-routing-profile",
	Short: "Describes the specified routing profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_describeRoutingProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_describeRoutingProfileCmd).Standalone()

		connect_describeRoutingProfileCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_describeRoutingProfileCmd.Flags().String("routing-profile-id", "", "The identifier of the routing profile.")
		connect_describeRoutingProfileCmd.MarkFlagRequired("instance-id")
		connect_describeRoutingProfileCmd.MarkFlagRequired("routing-profile-id")
	})
	connectCmd.AddCommand(connect_describeRoutingProfileCmd)
}
