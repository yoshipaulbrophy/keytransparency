// Copyright 2016 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package config has utilitites for loading configuration files from disk.
package config

import (
	"fmt"
	"log"

	"github.com/google/trillian"
	"github.com/google/trillian/client"
	"github.com/google/trillian/crypto/keys"
	"github.com/google/trillian/merkle/objhasher"
	"google.golang.org/grpc"
)

// LogClient creates a log client.
func LogClient(logID int64, logURL, pubKeyFile string) (client.VerifyingLogClient, error) {
	sthPubKey, err := keys.NewFromPublicPEMFile(pubKeyFile)
	if err != nil {
		log.Fatalf("Failed to open public key %v: %v", pubKeyFile, err)
	}
	// The log should be in a restricted access backend environment.
	cc, err := grpc.Dial(logURL, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to %v: %v", logURL, err)
	}
	log := client.New(logID, trillian.NewTrillianLogClient(cc),
		objhasher.ObjectHasher, sthPubKey)

	return log, nil
}
