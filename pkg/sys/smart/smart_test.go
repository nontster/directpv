//go:build !skip

// This file is part of MinIO DirectPV
// Copyright (c) 2021, 2022 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package smart

import (
	"fmt"
	"testing"
)

func TestSerialNumber(t1 *testing.T) {
	sn, err := GetSerialNumber("/dev/sdb")
	if err != nil {
		t1.Errorf("Cannot get serial number: %v", err)
	}
	fmt.Printf("\nSerial Number of /dev/sdb : %s", sn)
	sn, err = GetSerialNumber("/dev/nvme0n1")
	if err != nil {
		t1.Errorf("Cannot get serial number: %v", err)
	}
	fmt.Printf("\nSerial Number of /dev/nvme0n1: %s\n\n", sn)
}
