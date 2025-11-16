package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rtbfabric_createLinkCmd = &cobra.Command{
	Use:   "create-link",
	Short: "Creates a new link between gateways.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rtbfabric_createLinkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rtbfabric_createLinkCmd).Standalone()

		rtbfabric_createLinkCmd.Flags().String("attributes", "", "Attributes of the link.")
		rtbfabric_createLinkCmd.Flags().String("gateway-id", "", "The unique identifier of the gateway.")
		rtbfabric_createLinkCmd.Flags().Bool("http-responder-allowed", false, "Boolean to specify if an HTTP responder is allowed.")
		rtbfabric_createLinkCmd.Flags().String("log-settings", "", "Settings for the application logs.")
		rtbfabric_createLinkCmd.Flags().Bool("no-http-responder-allowed", false, "Boolean to specify if an HTTP responder is allowed.")
		rtbfabric_createLinkCmd.Flags().String("peer-gateway-id", "", "The unique identifier of the peer gateway.")
		rtbfabric_createLinkCmd.Flags().String("tags", "", "A map of the key-value pairs of the tag or tags to assign to the resource.")
		rtbfabric_createLinkCmd.MarkFlagRequired("gateway-id")
		rtbfabric_createLinkCmd.MarkFlagRequired("log-settings")
		rtbfabric_createLinkCmd.Flag("no-http-responder-allowed").Hidden = true
		rtbfabric_createLinkCmd.MarkFlagRequired("peer-gateway-id")
	})
	rtbfabricCmd.AddCommand(rtbfabric_createLinkCmd)
}
