// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package suricata

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "suricata", asset.ModuleFieldsPri, AssetSuricata); err != nil {
		panic(err)
	}
}

// AssetSuricata returns asset data.
// This is the base64 encoded gzipped contents of module/suricata.
func AssetSuricata() string {
	return "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"
}
