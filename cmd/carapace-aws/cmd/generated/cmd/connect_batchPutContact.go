package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_batchPutContactCmd = &cobra.Command{
	Use:   "batch-put-contact",
	Short: "Only the Amazon Connect outbound campaigns service principal is allowed to assume a role in your account and call this API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_batchPutContactCmd).Standalone()

	connect_batchPutContactCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	connect_batchPutContactCmd.Flags().String("contact-data-request-list", "", "List of individual contact requests.")
	connect_batchPutContactCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_batchPutContactCmd.MarkFlagRequired("contact-data-request-list")
	connect_batchPutContactCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_batchPutContactCmd)
}
