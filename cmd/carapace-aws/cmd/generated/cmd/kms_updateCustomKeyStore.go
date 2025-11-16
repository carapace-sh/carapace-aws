package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kms_updateCustomKeyStoreCmd = &cobra.Command{
	Use:   "update-custom-key-store",
	Short: "Changes the properties of a custom key store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kms_updateCustomKeyStoreCmd).Standalone()

	kms_updateCustomKeyStoreCmd.Flags().String("cloud-hsm-cluster-id", "", "Associates the custom key store with a related CloudHSM cluster.")
	kms_updateCustomKeyStoreCmd.Flags().String("custom-key-store-id", "", "Identifies the custom key store that you want to update.")
	kms_updateCustomKeyStoreCmd.Flags().String("key-store-password", "", "Enter the current password of the `kmsuser` crypto user (CU) in the CloudHSM cluster that is associated with the custom key store.")
	kms_updateCustomKeyStoreCmd.Flags().String("new-custom-key-store-name", "", "Changes the friendly name of the custom key store to the value that you specify.")
	kms_updateCustomKeyStoreCmd.Flags().String("xks-proxy-authentication-credential", "", "Changes the credentials that KMS uses to sign requests to the external key store proxy (XKS proxy).")
	kms_updateCustomKeyStoreCmd.Flags().String("xks-proxy-connectivity", "", "Changes the connectivity setting for the external key store.")
	kms_updateCustomKeyStoreCmd.Flags().String("xks-proxy-uri-endpoint", "", "Changes the URI endpoint that KMS uses to connect to your external key store proxy (XKS proxy).")
	kms_updateCustomKeyStoreCmd.Flags().String("xks-proxy-uri-path", "", "Changes the base path to the proxy APIs for this external key store.")
	kms_updateCustomKeyStoreCmd.Flags().String("xks-proxy-vpc-endpoint-service-name", "", "Changes the name that KMS uses to identify the Amazon VPC endpoint service for your external key store proxy (XKS proxy).")
	kms_updateCustomKeyStoreCmd.Flags().String("xks-proxy-vpc-endpoint-service-owner", "", "Changes the Amazon Web Services account ID that KMS uses to identify the Amazon VPC endpoint service for your external key store proxy (XKS proxy).")
	kms_updateCustomKeyStoreCmd.MarkFlagRequired("custom-key-store-id")
	kmsCmd.AddCommand(kms_updateCustomKeyStoreCmd)
}
