package secretsman

import (
	"context"
	"log"
	"os"
	"strings"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

type SecretManager struct {
	Context context.Context
	Client  *secretmanager.Client
	Logger  *log.Logger
	Secrets []string
}

func NewSecretManager(l *log.Logger) (*SecretManager, func() error) {
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		l.Fatal(err)
	}
	return &SecretManager{
		Context: ctx,
		Client:  client,
		Logger:  l,
		Secrets: make([]string, 0),
	}, client.Close
}

func (m *SecretManager) Setenv(secrets ...string) {
	for _, secret := range secrets {
		var secretResource strings.Builder
		secretResource.WriteString("projects/")
		secretResource.WriteString("holy-diver-297719")
		secretResource.WriteString("/secrets/")
		secretResource.WriteString(secret)
		secretResource.WriteString("/versions/")
		secretResource.WriteString("latest")

		accessSecretVersion := &secretmanagerpb.AccessSecretVersionRequest{
			Name: secretResource.String(),
		}

		resp, err := m.Client.AccessSecretVersion(m.Context, accessSecretVersion)
		if err != nil {
			m.Logger.Fatal(err)
		}
		secretValue := string(resp.Payload.GetData())
		os.Setenv(secret, secretValue)
	}
}
