package commands

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/open-policy-agent/conftest/policy"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const pullDesc = `
This command downloads individual policies from a remote location.

Several locations are supported by the pull command. Under the hood
conftest leverages go-getter (https://github.com/hashicorp/go-getter).
The  following protocols are supported for downloading policies:

	- OCI Registries
	- Local Files
	- Git
	- HTTP/HTTPS
	- Mercurial
	- Amazon S3
	- Google Cloud GCP

The location of the policies is specified by passing an URL, e.g.:

	$ conftest pull http://<my-policy-url>

Based on the protocol a different mechanism will be used to download the policy.
The pull command will also try to infer the protocol based on the URL if the 
URL does not contain a protocol. For example, the OCI mechanism will be used if
an azure registry URL is passed, e.g.

	$ conftest pull instrumenta.azurecr.io/my-registry


The policy location defaults to the policy directory in the local folder.
The location can be overridden with the '--policy' flag, e.g.:

	$ conftest push --policy <my-directory> <oci-url>
`

// NewPullCommand creates a new pull command to allow users
// to download individual policies
func NewPullCommand(ctx context.Context) *cobra.Command {
	cmd := cobra.Command{
		Use:   "pull <repository>",
		Short: "Download individual policies",
		Long:  pullDesc,
		Args:  cobra.MinimumNArgs(1),

		RunE: func(cmd *cobra.Command, args []string) error {
			policyDir := filepath.Join(".", viper.GetString("policy"))

			if err := policy.Download(ctx, policyDir, args); err != nil {
				return fmt.Errorf("download policies: %w", err)
			}

			return nil
		},
	}

	return &cmd
}
