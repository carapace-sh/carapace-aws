package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var signer_startSigningJobCmd = &cobra.Command{
	Use:   "start-signing-job",
	Short: "Initiates a signing job to be performed on the code provided.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(signer_startSigningJobCmd).Standalone()

	signer_startSigningJobCmd.Flags().String("client-request-token", "", "String that identifies the signing request.")
	signer_startSigningJobCmd.Flags().String("destination", "", "The S3 bucket in which to save your signed object.")
	signer_startSigningJobCmd.Flags().String("profile-name", "", "The name of the signing profile.")
	signer_startSigningJobCmd.Flags().String("profile-owner", "", "The AWS account ID of the signing profile owner.")
	signer_startSigningJobCmd.Flags().String("source", "", "The S3 bucket that contains the object to sign or a BLOB that contains your raw code.")
	signer_startSigningJobCmd.MarkFlagRequired("client-request-token")
	signer_startSigningJobCmd.MarkFlagRequired("destination")
	signer_startSigningJobCmd.MarkFlagRequired("profile-name")
	signer_startSigningJobCmd.MarkFlagRequired("source")
	signerCmd.AddCommand(signer_startSigningJobCmd)
}
