// Copyright (c) 2020 Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

// +build gofuzz

package fuzz

import (
	"github.com/mattermost/mattermost-server/mlog/human"
	"github.com/mattermost/mattermost-server/utils/markdown"
)

func FuzzParse(data []byte) int {
	_, _ = markdown.Parse(string(data))
	return 1
}

func FuzzParseLogMessage(data []byte) int {
	_ = human.ParseLogMessage(string(data))
	return 1
}
