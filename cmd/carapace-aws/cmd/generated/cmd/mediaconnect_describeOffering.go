package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_describeOfferingCmd = &cobra.Command{
	Use:   "describe-offering",
	Short: "Displays the details of an offering.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_describeOfferingCmd).Standalone()

	mediaconnect_describeOfferingCmd.Flags().String("offering-arn", "", "The ARN of the offering.")
	mediaconnect_describeOfferingCmd.MarkFlagRequired("offering-arn")
	mediaconnectCmd.AddCommand(mediaconnect_describeOfferingCmd)
}
