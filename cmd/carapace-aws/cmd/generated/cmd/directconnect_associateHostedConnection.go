package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_associateHostedConnectionCmd = &cobra.Command{
	Use:   "associate-hosted-connection",
	Short: "Associates a hosted connection and its virtual interfaces with a link aggregation group (LAG) or interconnect.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_associateHostedConnectionCmd).Standalone()

	directconnect_associateHostedConnectionCmd.Flags().String("connection-id", "", "The ID of the hosted connection.")
	directconnect_associateHostedConnectionCmd.Flags().String("parent-connection-id", "", "The ID of the interconnect or the LAG.")
	directconnect_associateHostedConnectionCmd.MarkFlagRequired("connection-id")
	directconnect_associateHostedConnectionCmd.MarkFlagRequired("parent-connection-id")
	directconnectCmd.AddCommand(directconnect_associateHostedConnectionCmd)
}
