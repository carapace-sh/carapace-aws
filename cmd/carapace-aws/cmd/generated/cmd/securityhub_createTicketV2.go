package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_createTicketV2Cmd = &cobra.Command{
	Use:   "create-ticket-v2",
	Short: "Grants permission to create a ticket in the chosen ITSM based on finding information for the provided finding metadata UID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_createTicketV2Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_createTicketV2Cmd).Standalone()

		securityhub_createTicketV2Cmd.Flags().String("client-token", "", "The client idempotency token.")
		securityhub_createTicketV2Cmd.Flags().String("connector-id", "", "The UUID of the connectorV2 to identify connectorV2 resource.")
		securityhub_createTicketV2Cmd.Flags().String("finding-metadata-uid", "", "The the unique ID for the finding.")
		securityhub_createTicketV2Cmd.MarkFlagRequired("connector-id")
		securityhub_createTicketV2Cmd.MarkFlagRequired("finding-metadata-uid")
	})
	securityhubCmd.AddCommand(securityhub_createTicketV2Cmd)
}
