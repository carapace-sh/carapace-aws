package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_createReservationCmd = &cobra.Command{
	Use:   "create-reservation",
	Short: "Creates an Amazon Redshift Serverless reservation, which gives you the option to commit to a specified number of Redshift Processing Units (RPUs) for a year at a discount from Serverless on-demand (OD) rates.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_createReservationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftServerless_createReservationCmd).Standalone()

		redshiftServerless_createReservationCmd.Flags().String("capacity", "", "The number of Redshift Processing Units (RPUs) to reserve.")
		redshiftServerless_createReservationCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		redshiftServerless_createReservationCmd.Flags().String("offering-id", "", "The ID of the offering associated with the reservation.")
		redshiftServerless_createReservationCmd.MarkFlagRequired("capacity")
		redshiftServerless_createReservationCmd.MarkFlagRequired("offering-id")
	})
	redshiftServerlessCmd.AddCommand(redshiftServerless_createReservationCmd)
}
