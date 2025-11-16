package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvert_disassociateCertificateCmd = &cobra.Command{
	Use:   "disassociate-certificate",
	Short: "Removes an association between the Amazon Resource Name (ARN) of an AWS Certificate Manager (ACM) certificate and an AWS Elemental MediaConvert resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvert_disassociateCertificateCmd).Standalone()

	mediaconvert_disassociateCertificateCmd.Flags().String("arn", "", "The ARN of the ACM certificate that you want to disassociate from your MediaConvert resource.")
	mediaconvert_disassociateCertificateCmd.MarkFlagRequired("arn")
	mediaconvertCmd.AddCommand(mediaconvert_disassociateCertificateCmd)
}
