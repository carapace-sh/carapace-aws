package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvert_associateCertificateCmd = &cobra.Command{
	Use:   "associate-certificate",
	Short: "Associates an AWS Certificate Manager (ACM) Amazon Resource Name (ARN) with AWS Elemental MediaConvert.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvert_associateCertificateCmd).Standalone()

	mediaconvert_associateCertificateCmd.Flags().String("arn", "", "The ARN of the ACM certificate that you want to associate with your MediaConvert resource.")
	mediaconvert_associateCertificateCmd.MarkFlagRequired("arn")
	mediaconvertCmd.AddCommand(mediaconvert_associateCertificateCmd)
}
