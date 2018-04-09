package service

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"time"

	spb "git.cafebazaar.ir/bar/saman-proto/proto-go"
	"github.com/cafebazaar/configman"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// SamanClientProvider Uses a grpc api to receive sorted reviews
type SamanClientProvider struct {
	config      configman.ConfigManager
	client      spb.SamanServiceClient
	conn        *grpc.ClientConn
	samanConfig *SamanClientConfig

	mutex sync.Mutex
}

type SamanClientConfig struct {
	enabled        bool
	debug          bool
	url            string
	certFile       string
	requestTimeout time.Duration
	dialTimeout    time.Duration
}

func NewSamanClientProvider(samanConfig configman.ConfigManager) (*SamanClientProvider, error) {
	smn := &SamanClientProvider{
		config: samanConfig,
	}
	err := smn.loadConfig()
	if err != nil {
		return nil, err
	}
	samanConfig.AddToChangeListeners(smn)
	return smn, nil
}

func (smn *SamanClientProvider) IsSamanEnabled() bool {
	return smn != nil && smn.samanConfig.enabled
}

func (smn *SamanClientProvider) DoesAppHaveInformativeReview(appId int32) (bool, error) {
	req := &spb.HasReviewBySortTypeRequest{
		SortType: "I",
		AppID:    appId,
	}

	ctx, cancel := context.WithTimeout(context.Background(), smn.samanConfig.requestTimeout)
	defer cancel()

	reply, err := smn.client.HasReviewBySortType(ctx, req)
	if err != nil {
		return false, err
	}

	return reply.GetResult(), nil
}

func (smn *SamanClientProvider) loadConfig() error {
	samanConfig := &SamanClientConfig{
		enabled:        smn.config.GetBool("saman.enabled"),
		debug:          smn.config.GetBool("saman.debug"),
		url:            smn.config.GetString("saman.url"),
		certFile:       smn.config.GetString("saman.cert-file"),
		requestTimeout: smn.config.GetDuration("saman.request-timeout"),
		dialTimeout:    smn.config.GetDuration("saman.dial-timeout"),
	}
	if !reflect.DeepEqual(samanConfig, smn.samanConfig) {
		smn.mutex.Lock()
		defer smn.mutex.Unlock()
		if samanConfig.debug {
			logrus.Debug("Preparing saman client connection")
		}

		if !samanConfig.enabled {
			logrus.WithField("url", samanConfig.url).
				Warning("saman service is disabled")
			return nil
		}
		smn.samanConfig = samanConfig

		creds, err := credentials.NewClientTLSFromFile(samanConfig.certFile, "localhost")
		if err != nil {
			fmt.Printf("failed to load tls from %s: %s", samanConfig.certFile, err)
			return err
		}

		grpcDialOpts := []grpc.DialOption{
			grpc.WithTransportCredentials(creds),
			grpc.WithTimeout(samanConfig.dialTimeout),
		}

		conn, err := grpc.Dial(samanConfig.url, grpcDialOpts...)
		if err != nil {
			logrus.Fatalf("could not connect to saman api (%s): %v", samanConfig.url, err)
		}
		smn.conn = conn
		smn.client = spb.NewSamanServiceClient(smn.conn)
		if samanConfig.debug {
			logrus.Debugf("New saman client is prepared %v", smn.samanConfig)
		}
	}
	return nil
}

func (smn *SamanClientProvider) OnConfigChanged() {
	err := smn.loadConfig()
	if err != nil {
		logrus.WithError(err).Error("Failed to reload config saman client provider")
	}
}

func (smn *SamanClientProvider) Close() error {
	smn.config.RemoveFromChangeListeners(smn)
	smn.conn.Close()
	return nil
}
