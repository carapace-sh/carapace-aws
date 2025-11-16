package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_updateIpsetCmd = &cobra.Command{
	Use:   "update-ipset",
	Short: "Updates the IPSet specified by the IPSet ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_updateIpsetCmd).Standalone()

	guardduty_updateIpsetCmd.Flags().Bool("activate", false, "The updated Boolean value that specifies whether the IPSet is active or not.")
	guardduty_updateIpsetCmd.Flags().String("detector-id", "", "The detectorID that specifies the GuardDuty service whose IPSet you want to update.")
	guardduty_updateIpsetCmd.Flags().String("expected-bucket-owner", "", "The Amazon Web Services account ID that owns the Amazon S3 bucket specified in the **location** parameter.")
	guardduty_updateIpsetCmd.Flags().String("ip-set-id", "", "The unique ID that specifies the IPSet that you want to update.")
	guardduty_updateIpsetCmd.Flags().String("location", "", "The updated URI of the file that contains the IPSet.")
	guardduty_updateIpsetCmd.Flags().String("name", "", "The unique ID that specifies the IPSet that you want to update.")
	guardduty_updateIpsetCmd.Flags().Bool("no-activate", false, "The updated Boolean value that specifies whether the IPSet is active or not.")
	guardduty_updateIpsetCmd.MarkFlagRequired("detector-id")
	guardduty_updateIpsetCmd.MarkFlagRequired("ip-set-id")
	guardduty_updateIpsetCmd.Flag("no-activate").Hidden = true
	guarddutyCmd.AddCommand(guardduty_updateIpsetCmd)
}
