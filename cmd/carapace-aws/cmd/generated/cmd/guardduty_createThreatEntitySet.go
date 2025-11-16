package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_createThreatEntitySetCmd = &cobra.Command{
	Use:   "create-threat-entity-set",
	Short: "Creates a new threat entity set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_createThreatEntitySetCmd).Standalone()

	guardduty_createThreatEntitySetCmd.Flags().Bool("activate", false, "A boolean value that indicates whether GuardDuty should start using the uploaded threat entity set to generate findings.")
	guardduty_createThreatEntitySetCmd.Flags().String("client-token", "", "The idempotency token for the create request.")
	guardduty_createThreatEntitySetCmd.Flags().String("detector-id", "", "The unique ID of the detector of the GuardDuty account for which you want to create a threat entity set.")
	guardduty_createThreatEntitySetCmd.Flags().String("expected-bucket-owner", "", "The Amazon Web Services account ID that owns the Amazon S3 bucket specified in the **location** parameter.")
	guardduty_createThreatEntitySetCmd.Flags().String("format", "", "The format of the file that contains the threat entity set.")
	guardduty_createThreatEntitySetCmd.Flags().String("location", "", "The URI of the file that contains the threat entity set.")
	guardduty_createThreatEntitySetCmd.Flags().String("name", "", "A user-friendly name to identify the threat entity set.")
	guardduty_createThreatEntitySetCmd.Flags().Bool("no-activate", false, "A boolean value that indicates whether GuardDuty should start using the uploaded threat entity set to generate findings.")
	guardduty_createThreatEntitySetCmd.Flags().String("tags", "", "The tags to be added to a new threat entity set resource.")
	guardduty_createThreatEntitySetCmd.MarkFlagRequired("activate")
	guardduty_createThreatEntitySetCmd.MarkFlagRequired("detector-id")
	guardduty_createThreatEntitySetCmd.MarkFlagRequired("format")
	guardduty_createThreatEntitySetCmd.MarkFlagRequired("location")
	guardduty_createThreatEntitySetCmd.MarkFlagRequired("name")
	guardduty_createThreatEntitySetCmd.Flag("no-activate").Hidden = true
	guardduty_createThreatEntitySetCmd.MarkFlagRequired("no-activate")
	guarddutyCmd.AddCommand(guardduty_createThreatEntitySetCmd)
}
